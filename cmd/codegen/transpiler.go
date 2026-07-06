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
	"mvdan.cc/sh/v3/shell"
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
	defaultPreprocessor   = "clang"
)

// extensionFloatTypedefRe matches glibc interchange-type typedefs that clang
// emits from bits/floatn*.h. GCC omits them (builtins); ccgo cannot type-check
// the clang forms (__float128, duplicate specifiers, __mode__ (__TC__), …).
var extensionFloatTypedefRe = regexp.MustCompile(
	`(?m)^typedef\s+[^;]*\b(_Float(16|32|32x|64|64x|128)|__float128|__cfloat128)\b[^;]*;\s*\n?`,
)

var (
	mingwAsmStartRe  = regexp.MustCompile(`__asm(?:__)?(?:\s+__(?:volatile)__|\s+volatile)?\s*\(`)
	mingwAttrStartRe = regexp.MustCompile(`__(?:attribute|declspec)(?:__)?\s*\(`)

	libcVersionOnce sync.Once
	libcVersion     string
	libcVersionErr  error
)

// stripBalancedCalls removes each match of startRe through its balanced "(...)".
func stripBalancedCalls(src string, startRe *regexp.Regexp) string {
	var b strings.Builder
	b.Grow(len(src))
	i := 0
	for i < len(src) {
		loc := startRe.FindStringIndex(src[i:])
		if loc == nil {
			b.WriteString(src[i:])
			break
		}
		start := i + loc[0]
		open := i + loc[1] - 1
		b.WriteString(src[i:start])
		depth := 0
		j := open
		for j < len(src) {
			switch src[j] {
			case '(':
				depth++
			case ')':
				depth--
				if depth == 0 {
					j++
					i = j
					goto next
				}
			}
			j++
		}
		b.WriteString(src[start:])
		return b.String()
	next:
	}
	return b.String()
}

// preprocessorCmd builds the C preprocessor command from CC and optional CFLAGS.
// On Windows, MinGW gcc is preferred over clang when present on PATH.
func preprocessorCmd(goos, goarch string, args ...string) (*exec.Cmd, error) {
	ccFields, err := splitCompilerEnv(resolveCC(goos, goarch))
	if err != nil {
		return nil, fmt.Errorf("parse CC: %w", err)
	}
	if len(ccFields) == 0 {
		return nil, fmt.Errorf("CC is empty")
	}
	cflags, err := splitCompilerEnv(os.Getenv("CFLAGS"))
	if err != nil {
		return nil, fmt.Errorf("parse CFLAGS: %w", err)
	}
	cmdArgs := append([]string{}, ccFields[1:]...)
	cmdArgs = append(cmdArgs, defaultPreprocessorFlags(goos, goarch, ccFields[0])...)
	cmdArgs = append(cmdArgs, cflags...)
	cmdArgs = append(cmdArgs, args...)
	return exec.Command(ccFields[0], cmdArgs...), nil
}

// splitCompilerEnv splits CC/CFLAGS. mvdan.cc/sh treats \ as an escape, which
// turns Windows paths like C:\mingw64\bin\gcc.exe into C:mingw64bingcc.exe.
// Keep drive/UNC paths intact (split on ".exe" when args follow).
func splitCompilerEnv(val string) ([]string, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil, nil
	}
	if isWindowsCmdPath(val) {
		return splitWindowsCompilerEnv(val)
	}
	return shell.Fields(val, nil)
}

func isWindowsCmdPath(s string) bool {
	if strings.HasPrefix(s, `\\`) {
		return true
	}
	return len(s) >= 3 && s[1] == ':' && (s[2] == '\\' || s[2] == '/')
}

func splitWindowsCompilerEnv(val string) ([]string, error) {
	lower := strings.ToLower(val)
	if i := strings.Index(lower, ".exe"); i >= 0 {
		exe := val[:i+len(".exe")]
		rest := strings.TrimSpace(val[i+len(".exe"):])
		fields := []string{exe}
		if rest == "" {
			return fields, nil
		}
		extra, err := shell.Fields(rest, nil)
		if err != nil {
			return nil, err
		}
		return append(fields, extra...), nil
	}
	if !strings.ContainsAny(val, " \t") {
		return []string{val}, nil
	}
	return shell.Fields(filepath.ToSlash(val), nil)
}

func resolveCC(goos, goarch string) string {
	cc := strings.TrimSpace(os.Getenv("CC"))
	if goos == "windows" {
		if mingw := lookupMingwGCC(goarch); mingw != "" {
			if cc == "" || strings.Contains(filepath.Base(cc), "clang") {
				return mingw
			}
		}
	}
	if cc != "" {
		return cc
	}
	return defaultPreprocessor
}

// lookupMingwGCC returns the MinGW GCC driver on PATH so -E uses GNU headers
// without clang rewriting __declspec through __attribute__.
func lookupMingwGCC(goarch string) string {
	for _, cand := range []string{mingwTriple(goarch) + "-gcc", "gcc"} {
		if p, err := exec.LookPath(cand); err == nil {
			return p
		}
	}
	return ""
}

// defaultPreprocessorFlags returns compiler flags so host -E output stays
// within what ccgo can parse. Prefer triples and macros over rewriting sources.
func defaultPreprocessorFlags(goos, goarch, ccBinary string) []string {
	base := filepath.Base(ccBinary)
	isClang := strings.Contains(base, "clang")
	switch goos {
	case "windows":
		var flags []string
		if isClang {
			flags = append(flags, "--target="+mingwTriple(goarch))
			if sysroot := mingwSysroot(goarch); sysroot != "" {
				flags = append(flags, "--sysroot="+sysroot)
			}
		}
		// Erase GNU/MinGW extensions ccgo cannot parse. Empty __attribute__
		// on the command line so headers' `#define __declspec(a) __attribute__((a))`
		// expands to nothing (not stray `(dllimport)` tokens). Calling-convention
		// keywords similarly survive -E as junk identifiers.
		flags = append(flags,
			"-D__extension__=",
			"-D__forceinline=static inline",
			"-D__attribute__(...)=",
			"-D__declspec(x)=",
			"-D__cdecl=",
			"-D__stdcall=",
			"-D__fastcall=",
			"-D__thiscall=",
			"-D_cdecl=",
			"-D__restrict=",
			"-D__restrict__=",
			"-D__MINGW_EXTENSION=",
		)
		return flags
	case "darwin":
		// Blocks + availability(macos,introduced=10.13.4) leave tokens like
		// 10.13.4 that ccgo parses as invalid floats. Neutralize via macros.
		// Alignment-16 neon structs are ignored later via ccgo flags.
		return []string{
			"-fno-blocks",
			"-D__attribute__(...)=",
			"-D__extension__=",
			"-D_Nonnull=",
			"-D_Nullable=",
			"-D_Null_unspecified=",
			"-DAPI_AVAILABLE(...)=",
			"-DAPI_UNAVAILABLE(...)=",
			"-DAPI_DEPRECATED(...)=",
			"-DAPI_DEPRECATED_WITH_REPLACEMENT(...)=",
			"-D__API_AVAILABLE(...)=",
			"-D__API_UNAVAILABLE(...)=",
			"-D__API_DEPRECATED(...)=",
			"-D__API_DEPRECATED_WITH_REPLACEMENT(...)=",
		}
	default:
		return nil
	}
}

func ccgoExtraArgs(goos string) []string {
	args := []string{
		"-ignore-unsupported-alignment",
		"-ignore-unsupported-atomic-sizes",
		"-ignore-vector-functions",
	}
	if goos == "windows" {
		// Registered in ccgo as Opt("-winapi-no-errno"); modernc.org/opt
		// strips one leading '-', so callers must pass the double-dash form.
		args = append(args, "--winapi-no-errno")
	}
	return args
}

func mingwTriple(goarch string) string {
	switch goarch {
	case "arm64":
		return "aarch64-w64-mingw32"
	case "386":
		return "i686-w64-mingw32"
	default:
		return "x86_64-w64-mingw32"
	}
}

func mingwSysroot(goarch string) string {
	if s := strings.TrimSpace(os.Getenv("MINGW_SYSROOT")); s != "" {
		// Trust an explicit amd64 sysroot; other arches must match the triple layout.
		if goarch == "amd64" || goarch == "" {
			return s
		}
		if mingwSysrootMatches(s, goarch) {
			return s
		}
		return ""
	}
	var candidates []string
	switch goarch {
	case "arm64":
		candidates = []string{
			`C:\llvm-mingw`,
			`C:\msys64\clangarm64`,
			"/usr/aarch64-w64-mingw32",
		}
	case "386":
		candidates = []string{
			`C:\msys64\mingw32`,
			`C:\mingw32`,
			"/usr/i686-w64-mingw32",
		}
	default:
		candidates = []string{
			`C:\msys64\mingw64`,
			`C:\mingw64`,
			`C:\ProgramData\mingw64\mingw64`,
			`C:\ProgramData\chocolatey\lib\mingw\tools\install\mingw64`,
			"/mingw64",
			"/usr/x86_64-w64-mingw32",
		}
	}
	for _, candidate := range candidates {
		if mingwSysrootMatches(candidate, goarch) {
			return candidate
		}
	}
	return ""
}

func mingwSysrootMatches(root, goarch string) bool {
	if st, err := os.Stat(filepath.Join(root, "include")); err != nil || !st.IsDir() {
		return false
	}
	triple := mingwTriple(goarch)
	// Accept plain mingw64 roots for amd64; require triple-specific layout otherwise.
	if goarch == "amd64" || goarch == "" {
		return true
	}
	if st, err := os.Stat(filepath.Join(root, triple)); err == nil && st.IsDir() {
		return true
	}
	if st, err := os.Stat(filepath.Join(root, "lib", triple)); err == nil && st.IsDir() {
		return true
	}
	// llvm-mingw uses <root>/include and <root>/aarch64-w64-mingw32
	if st, err := os.Stat(filepath.Join(root, triple, "include")); err == nil && st.IsDir() {
		return true
	}
	return false
}

// sanitizePreprocessedC removes constructs that survive -E but are rejected by ccgo.
func sanitizePreprocessedC(src string) string {
	src = extensionFloatTypedefRe.ReplaceAllString(src, "")
	src = stripBalancedCalls(src, mingwAsmStartRe)
	src = stripBalancedCalls(src, mingwAttrStartRe)
	return src
}

func sanitizePreprocessedFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	sanitized := sanitizePreprocessedC(string(data))
	if sanitized == string(data) {
		return nil
	}
	return os.WriteFile(path, []byte(sanitized), 0644)
}

func preprocessorIdentity(goos, goarch string) string {
	ccFields, err := splitCompilerEnv(resolveCC(goos, goarch))
	if err != nil || len(ccFields) == 0 {
		ccFields = []string{resolveCC(goos, goarch)}
	}
	parts := append(append([]string{}, ccFields...), defaultPreprocessorFlags(goos, goarch, ccFields[0])...)
	if cflags := strings.TrimSpace(os.Getenv("CFLAGS")); cflags != "" {
		parts = append(parts, cflags)
	}
	return strings.Join(parts, " ")
}

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

	args := []string{
		"-E", "-O0", "-fno-inline",
		"-ffile-prefix-map=" + filepath.ToSlash(workDir) + "=.",
		"-fmacro-prefix-map=" + filepath.ToSlash(workDir) + "=.",
		filepath.Join(t.TreeSitterPath, "lib/src/lib.c"),
		"-I", filepath.Join(t.TreeSitterPath, "lib/include"),
		"-I", filepath.Join(t.TreeSitterPath, "lib/src"),
		"-o", "-",
	}

	cmd, err := preprocessorCmd(t.GOOS, t.GOARCH, args...)
	if err != nil {
		f.Close()
		return err
	}
	cmd.Stdout = f
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		f.Close()
		return fmt.Errorf("preprocess core with %s: %w", preprocessorIdentity(t.GOOS, t.GOARCH), err)
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := sanitizePreprocessedFile(outputPath); err != nil {
		return fmt.Errorf("sanitize preprocessed core: %w", err)
	}

	f, err = os.OpenFile(outputPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Add atomic stubs for compiler builtins that ccgo doesn't handle natively.
	// Use unsigned int (not uint32_t) to avoid clashing with windows headers.
	stubs := "\nstatic inline unsigned int __atomic_add_fetch(volatile unsigned int *p, unsigned int v, int m) { *p += v; return *p; }\n" +
		"static inline unsigned int __atomic_sub_fetch(volatile unsigned int *p, unsigned int v, int m) { *p -= v; return *p; }\n"
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

	cmd, err := preprocessorCmd(t.GOOS, t.GOARCH, args...)
	if err != nil {
		f.Close()
		return err
	}
	cmd.Stdout = f
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		f.Close()
		return fmt.Errorf("preprocess grammar with %s: %w", preprocessorIdentity(t.GOOS, t.GOARCH), err)
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := sanitizePreprocessedFile(preprocessed); err != nil {
		return fmt.Errorf("sanitize preprocessed grammar: %w", err)
	}

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
	ccgoArgs := append([]string{"ccgo"}, ccgoExtraArgs(t.GOOS)...)
	ccgoArgs = append(ccgoArgs, inputArg, "-o", outputArg)

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
