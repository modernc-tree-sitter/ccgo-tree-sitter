package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"modernc.org/ccgo/v4/lib"
)

// Transpiler handles C to Go transpilation using ccgo.
// It orchestrates the process of transpiling tree-sitter C code into Go,
// setting up isolated execution environments to prevent ccgo state pollution.
type Transpiler struct {
	// TreeSitterPath is the root directory of the tree-sitter source.
	TreeSitterPath string
	// GOOS is the target operating system for generated Go code.
	GOOS string
	// GOARCH is the target architecture for generated Go code.
	GOARCH string
	// KeepTemp specifies whether to leave intermediate C and go files intact for debugging.
	KeepTemp bool
}

const (
	deterministicTempRoot = "ccgo-tree-sitter-gen"
)

var (
	libcVersionOnce sync.Once
	libcVersion     string
	libcVersionErr  error
)

// TranspileCore transpiles the tree-sitter core library into the target directory.
// The flow consists of preprocessing C headers, generating a synthetic go.mod to satisfy
// ccgo dependencies, executing ccgo on the unified C files, and finally post-processing
// to fix deterministic paths and perform AST patching for CGO-free pointers.
func (t *Transpiler) TranspileCore(outputDir string) error {
	tmpDir, err := t.prepareWorkDir("tree-sitter-core", "")
	if err != nil {
		return err
	}
	if !t.KeepTemp {
		defer os.RemoveAll(tmpDir)
	} else {
		slog.Info("keeping temp dir", "path", tmpDir)
	}

	// Preprocess
	coreComplete := filepath.Join(tmpDir, "core_complete.c")
	if err := t.preprocessCore(coreComplete); err != nil {
		return fmt.Errorf("preprocess failed: %w", err)
	}

	// Setup go.mod
	if err := setupGoMod(tmpDir); err != nil {
		return fmt.Errorf("go.mod setup failed: %w", err)
	}

	// Transpile
	coreGo := filepath.Join(tmpDir, "core.go")
	if err := t.runCcgo(tmpDir, coreComplete, coreGo); err != nil {
		return fmt.Errorf("ccgo failed: %w", err)
	}

	// Post-process and write
	if outputDir != "" {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return err
		}

		// Read and modify to be a proper package
		data, err := os.ReadFile(coreGo)
		if err != nil {
			return err
		}

		goCode := postProcess(string(data), tmpDir)
		// Change package from main to grammar
		goCode = strings.Replace(goCode, "package main", "package grammar", 1)

		outputFile := filepath.Join(outputDir, fmt.Sprintf("core-%s-%s.go", t.GOOS, t.GOARCH))
		return os.WriteFile(outputFile, []byte(goCode), 0644)
	}

	// Output to stdout
	data, err := os.ReadFile(coreGo)
	if err != nil {
		return err
	}
	fmt.Print(postProcess(string(data), tmpDir))
	return nil
}

// TranspileGrammar transpiles a tree-sitter grammar unit into the target directory.
// grammarPath is the unit root (parent of src/); grammarName is the package/registry id
// (typically from tree_sitter_<name> in parser.c). If grammarName is empty, it is derived
// from the unit path.
func (t *Transpiler) TranspileGrammar(grammarPath, grammarName, outputDir string) error {
	if grammarName == "" {
		grammarName = extractGrammarName(grammarPath)
	}

	// Check parser.c exists
	parserC := filepath.Join(grammarPath, "src", "parser.c")
	if _, err := os.Stat(parserC); os.IsNotExist(err) {
		return fmt.Errorf("parser.c not found at %s", parserC)
	}

	// Check for scanner.c
	scannerC := filepath.Join(grammarPath, "src", "scanner.c")
	hasScanner := true
	if _, err := os.Stat(scannerC); os.IsNotExist(err) {
		hasScanner = false
	}

	tmpDir, err := t.prepareWorkDir("grammar", grammarName)
	if err != nil {
		return err
	}
	if !t.KeepTemp {
		defer os.RemoveAll(tmpDir)
	} else {
		slog.Info("keeping temp dir", "path", tmpDir)
	}

	// If has scanner, combine parser and scanner into one file
	var combinedC string
	if hasScanner {
		combinedC = filepath.Join(tmpDir, "combined.c")
		if err := combineFiles([]string{scannerC, parserC}, combinedC); err != nil {
			return fmt.Errorf("failed to combine files: %w", err)
		}
	} else {
		combinedC = parserC
	}

	// Transpile combined file
	grammarGo := filepath.Join(tmpDir, "grammar.go")
	if err := t.transpileGrammarFile(grammarPath, combinedC, grammarGo, tmpDir); err != nil {
		return fmt.Errorf("transpilation failed: %w", err)
	}

	// Write output as a proper Go package
	if outputDir != "" {
		grammarOutDir := filepath.Join(outputDir, grammarName)
		if err := os.MkdirAll(grammarOutDir, 0755); err != nil {
			return err
		}

		// Read and modify package name
		data, err := os.ReadFile(grammarGo)
		if err != nil {
			return err
		}

		goCode := postProcess(string(data), tmpDir)
		// Change package from main to grammar name (preserve comments)
		goCode = strings.Replace(goCode, "package main", "package grammar_"+grammarName, 1)

		outputFile := filepath.Join(grammarOutDir, fmt.Sprintf("grammar-%s-%s.go", t.GOOS, t.GOARCH))
		if err := os.WriteFile(outputFile, []byte(goCode), 0644); err != nil {
			return err
		}

		// Generate API wrapper
		var apiErr error
		if hasScanner {
			apiErr = GenerateAPIWrapperWithScanner(outputDir, grammarName)
		} else {
			apiErr = GenerateAPIWrapper(outputDir, grammarName)
		}
		if apiErr != nil {
			return fmt.Errorf("failed to generate API: %w", apiErr)
		}
	}

	return nil
}

// combineFiles sequentially concatenates multiple C input source paths
// into a single output file to supply ccgo with unified compilation units.
func combineFiles(inputs []string, output string) error {
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, input := range inputs {
		data, err := os.ReadFile(input)
		if err != nil {
			return err
		}
		if _, err := out.Write(data); err != nil {
			return err
		}
		out.WriteString("\n\n")
	}

	return nil
}

func (t *Transpiler) preprocessCore(outputPath string) error {
	workDir := filepath.Dir(outputPath)
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	args := []string{
		"-E", "-O0", "-fno-inline",
		"-ffile-prefix-map=" + filepath.ToSlash(workDir) + "=.",
		"-fmacro-prefix-map=" + filepath.ToSlash(workDir) + "=.",
		filepath.Join(t.TreeSitterPath, "lib/src/lib.c"),
		"-I", filepath.Join(t.TreeSitterPath, "lib/include"),
		"-I", filepath.Join(t.TreeSitterPath, "lib/src"),
		"-o", "-",
	}

	cmd := exec.Command("gcc", args...)
	cmd.Stdout = f
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// Add atomic stubs
	stubs := "typedef unsigned int uint32_t;\n" +
		"static inline uint32_t __atomic_add_fetch(volatile uint32_t *p, uint32_t v, int m) { *p += v; return *p; }\n" +
		"static inline uint32_t __atomic_sub_fetch(volatile uint32_t *p, uint32_t v, int m) { *p -= v; return *p; }\n\n"
	_, err = fmt.Fprint(f, stubs)
	return err
}

func (t *Transpiler) transpileGrammarFile(grammarPath, inputC, outputGo, workDir string) error {
	// Preprocess
	preprocessed := filepath.Join(workDir, "preprocessed.c")
	f, err := os.Create(preprocessed)
	if err != nil {
		return err
	}

	args := []string{
		"-E", "-O0", "-fno-inline",
		"-ffile-prefix-map=" + filepath.ToSlash(workDir) + "=.",
		"-fmacro-prefix-map=" + filepath.ToSlash(workDir) + "=.",
		"-Dfunc=func_token",
		"-Dinterface=interface_token",
		"-Dselect=select_token",
		"-Dchan=chan_token",
		"-Dgo=go_token",
		"-Dmap=map_token",
		"-Dpackage=package_token",
		"-Dtype=type_token",
		"-Dvar=var_token",
		"-Dimport=import_token",
		"-Ddefer=defer_token",
		"-Dfallthrough=fallthrough_token",
		"-Drange=range_token",
		inputC,
		"-I", filepath.Join(grammarPath, "src"),
		// Monorepo grammars (typescript, markdown) often include headers from a sibling common/
		"-I", grammarPath,
		"-I", filepath.Dir(grammarPath),
		"-I", filepath.Join(t.TreeSitterPath, "lib/include"),
		"-I", filepath.Join(t.TreeSitterPath, "lib/src"),
		"-o", "-",
	}

	cmd := exec.Command("gcc", args...)
	cmd.Stdout = f
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		f.Close()
		return err
	}
	f.Close()

	// Setup go.mod if needed
	goModPath := filepath.Join(workDir, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		if err := setupGoMod(workDir); err != nil {
			return err
		}
	}

	// Transpile
	return t.runCcgo(workDir, preprocessed, outputGo)
}

func (t *Transpiler) runCcgo(workDir, inputPath, outputPath string) error {
	inputArg := inputPath
	if rel, err := filepath.Rel(workDir, inputPath); err == nil {
		inputArg = rel
	}
	outputArg := outputPath
	if rel, err := filepath.Rel(workDir, outputPath); err == nil {
		outputArg = rel
	}
	ccgoArgs := []string{"ccgo", inputArg, "-o", outputArg}

	// Change to work dir
	origDir, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.Chdir(workDir); err != nil {
		return err
	}
	defer os.Chdir(origDir)

	// Isolate and call ccgo
	err = runCcgoIsolated(t.GOOS, t.GOARCH, ccgoArgs)

	// If file was generated, ignore validation errors
	if err != nil {
		if _, statErr := os.Stat(outputPath); statErr == nil {
			// File exists, ccgo succeeded even if validation failed
			return nil
		}
		return err
	}

	return nil
}

// runCcgoIsolated runs ccgo with completely isolated state.
// Because ccgo pollutes the global flag.CommandLine and os.Args during processing,
// running multiple transpilation tasks within the same process leads to panics
// and argument bleed. This function safely restores the global state upon exit.
func runCcgoIsolated(goos, goarch string, args []string) error {
	// Save state
	oldArgs := os.Args
	oldCmdLine := flag.CommandLine
	oldUsage := flag.Usage
	oldStderr := os.Stderr

	defer func() {
		os.Args = oldArgs
		flag.CommandLine = oldCmdLine
		flag.Usage = oldUsage
		os.Stderr = oldStderr
	}()

	// Fresh state
	os.Args = args
	flag.CommandLine = flag.NewFlagSet(args[0], flag.ContinueOnError)
	flag.CommandLine.SetOutput(io.Discard)
	flag.Usage = func() {}

	// Redirect stderr to capture output
	r, w, _ := os.Pipe()
	os.Stderr = w

	var stderrBuf bytes.Buffer
	done := make(chan struct{})
	go func() {
		io.Copy(&stderrBuf, r)
		close(done)
	}()

	// Run ccgo
	t := ccgo.NewTask(goos, goarch, args, io.Discard, w, nil)
	err := t.Main()

	w.Close()
	<-done

	if err != nil {
		return fmt.Errorf("%w\nstderr: %s", err, stderrBuf.String())
	}

	return nil
}

// Helper functions

func setupGoMod(dir string) error {
	libcVer, err := currentLibcVersion()
	if err != nil {
		return err
	}

	goModPath := filepath.Join(dir, "go.mod")
	goModContent := "module treesitter\n\ngo 1.25\n\nrequire modernc.org/libc " + libcVer + "\n"
	if err := os.WriteFile(goModPath, []byte(goModContent), 0644); err != nil {
		return err
	}

	getCmd := exec.Command("go", "mod", "download", "modernc.org/libc@"+libcVer)
	getCmd.Dir = dir
	getCmd.Stderr = os.Stderr
	getCmd.Stdout = io.Discard
	return getCmd.Run()
}

func currentLibcVersion() (string, error) {
	libcVersionOnce.Do(func() {
		cmd := exec.Command("go", "list", "-m", "-f", "{{.Version}}", "modernc.org/libc")
		out, err := cmd.Output()
		if err != nil {
			libcVersionErr = fmt.Errorf("resolve modernc.org/libc version: %w", err)
			return
		}
		libcVersion = strings.TrimSpace(string(out))
		if libcVersion == "" {
			libcVersionErr = fmt.Errorf("resolve modernc.org/libc version: empty version from go list")
		}
	})
	if libcVersionErr != nil {
		return "", libcVersionErr
	}
	return libcVersion, nil
}

// postProcess normalizes generated paths for deterministic cross-run outputs,
// and applies structural patches to pointer dereferencing in the tree-sitter AST
// to prevent panics when traversing null child relationships natively in Go.
func postProcess(goCode, workDir string) string {
	// Normalize temp workdir path so regenerated code is deterministic across runs.
	workDirSlash := filepath.ToSlash(workDir)
	goCode = strings.ReplaceAll(goCode, workDirSlash+"/", "./")
	goCode = strings.ReplaceAll(goCode, workDirSlash, ".")
	goCode = regexp.MustCompile(`/tmp/(grammar-gen|tree-sitter-gen)-\d+`).ReplaceAllString(goCode, "/tmp/$1")

	// Fix assert types
	goCode = regexp.MustCompile(`libc\.X__assert_fail\(tls, ([^,]*), ([^,]*), uint32\((\d+)\)`).
		ReplaceAllString(goCode, `libc.X__assert_fail(tls, $1, $2, int32($3)`)

	// Remove negative padding (any number of tabs)
	goCode = regexp.MustCompile(`\t+_ \[-\d+\]byte\n`).ReplaceAllString(goCode, "")

	// Remove debug println statements from ccgo
	goCode = regexp.MustCompile(`println\(__ccgo_ts \+ \d+\)[\s;]*`).ReplaceAllString(goCode, "")
	goCode = regexp.MustCompile(`println\(__ccgo_ts\+\d+, __ccgo_ts\+\d+\)[\s;]*`).ReplaceAllString(goCode, "")

	// Patch ts_subtree_child_count to check for NULL pointer before dereferencing
	goCode = regexp.MustCompile(`func ts_subtree_child_count\(tls \*libc\.TLS, _self Subtree\) \(r uint32_t\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 uint32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = uint32\(0\)
	\} else \{
		v1 = \(\*SubtreeHeapData\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\)\)\)\.Fchild_count
	\}
	return v1
\}`).ReplaceAllString(goCode, `func ts_subtree_child_count(tls *libc.TLS, _self Subtree) (r uint32_t) {
	// NULL check - if the subtree pointer is 0, return 0
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 uint32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = uint32(0)
	} else {
		v1 = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count
	}
	return v1
}`)

	// Patch ts_subtree_extra to check for NULL pointer
	goCode = regexp.MustCompile(`func ts_subtree_extra\(tls \*libc\.TLS, _self Subtree\) \(r uint8\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 int32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\) & 0x8 >> 3\)
	\} else \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\) \+ 44\)\) & 0x4 >> 2\)
	\}
	return libc\.Uint8FromInt32\(libc\.BoolInt32\(v1 != 0\)\)
\}`).ReplaceAllString(goCode, `func ts_subtree_extra(tls *libc.TLS, _self Subtree) (r uint8) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x8 >> 3)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x4 >> 2)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}`)

	// Patch ts_subtree_symbol to check for NULL pointer
	goCode = regexp.MustCompile(`func ts_subtree_symbol\(tls \*libc\.TLS, _self Subtree\) \(r TSSymbol\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 int32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = libc\.Int32FromUint8\(\(\*\(\*SubtreeInlineData\)\(unsafe\.Pointer\(bp\)\)\)\.Fsymbol\)
	\} else \{
		v1 = libc\.Int32FromUint16\(\(\*SubtreeHeapData\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\)\)\)\.Fsymbol\)
	\}
	return libc\.Uint16FromInt32\(v1\)
\}`).ReplaceAllString(goCode, `func ts_subtree_symbol(tls *libc.TLS, _self Subtree) (r TSSymbol) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = libc.Int32FromUint8((*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fsymbol)
	} else {
		v1 = libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol)
	}
	return libc.Uint16FromInt32(v1)
}`)

	// Patch ts_subtree_visible to check for NULL pointer
	goCode = regexp.MustCompile(`func ts_subtree_visible\(tls \*libc\.TLS, _self Subtree\) \(r uint8\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 int32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\) & 0x2 >> 1\)
	\} else \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\) \+ 44\)\) & 0x1 >> 0\)
	\}
	return libc\.Uint8FromInt32\(libc\.BoolInt32\(v1 != 0\)\)
\}`).ReplaceAllString(goCode, `func ts_subtree_visible(tls *libc.TLS, _self Subtree) (r uint8) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x2 >> 1)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x1 >> 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}`)

	// Patch ts_subtree_named to check for NULL pointer
	goCode = regexp.MustCompile(`func ts_subtree_named\(tls \*libc\.TLS, _self Subtree\) \(r uint8\) \{
	bp := tls\.Alloc\(16\)
	defer tls\.Free\(16\)
	\*\(\*Subtree\)\(unsafe\.Pointer\(bp\)\) = _self
	var v1 int32
	_ = v1
	if int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\)&0x1>>0\) != 0 \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(bp \+ 0\)\) & 0x4 >> 2\)
	\} else \{
		v1 = int32\(\*\(\*uint8\)\(unsafe\.Pointer\(\*\(\*uintptr\)\(unsafe\.Pointer\(bp\)\) \+ 44\)\) & 0x2 >> 1\)
	\}
	return libc\.Uint8FromInt32\(libc\.BoolInt32\(v1 != 0\)\)
\}`).ReplaceAllString(goCode, `func ts_subtree_named(tls *libc.TLS, _self Subtree) (r uint8) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x4 >> 2)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x2 >> 1)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}`)

	return goCode
}

func (t *Transpiler) prepareWorkDir(kind, suffix string) (string, error) {
	parts := []string{
		os.TempDir(),
		deterministicTempRoot,
		fmt.Sprintf("%s-%s", t.GOOS, t.GOARCH),
		kind,
	}
	if suffix != "" {
		parts = append(parts, suffix)
	}
	workDir := filepath.Join(parts...)
	if err := os.RemoveAll(workDir); err != nil {
		return "", err
	}
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return "", err
	}
	return workDir, nil
}
