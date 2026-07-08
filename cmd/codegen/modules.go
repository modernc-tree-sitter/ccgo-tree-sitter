package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/mod/modfile"
)

const (
	rootModulePath    = "github.com/lucasew/ccgo-tree-sitter"
	grammarModulePath = rootModulePath + "/grammar"
	moduleGoVersion   = "1.25.0"
	// Keep in sync with root go.mod require modernc.org/libc.
	libcModuleVersion = "v1.67.6"
	localPseudoVer    = "v0.0.0"
)

// ensureGrammarModules writes grammar/go.mod, grammar/<lang>/go.mod (with local
// replace directives), go.work, and wires root go.mod require+replace entries.
func ensureGrammarModules(outputDir string) error {
	abs, err := filepath.Abs(outputDir)
	if err != nil {
		return err
	}
	grammarDir := filepath.Join(abs, "grammar")
	if err := os.MkdirAll(grammarDir, 0755); err != nil {
		return err
	}
	if err := writeCoreGoMod(grammarDir); err != nil {
		return err
	}

	langs, err := listGrammarLangs(grammarDir)
	if err != nil {
		return err
	}
	for _, lang := range langs {
		if err := writeLangGoMod(grammarDir, lang); err != nil {
			return err
		}
	}
	if err := writeGoWork(abs, langs); err != nil {
		return err
	}
	if err := updateRootGoMod(abs, langs); err != nil {
		return fmt.Errorf("update root go.mod: %w", err)
	}
	slog.Info("wrote grammar modules", "core", 1, "langs", len(langs), "go.work", true)
	return nil
}

// tidyGrammarModules runs go work sync and go mod tidy for root, core, and each
// language module so go.sum files exist.
func tidyGrammarModules(outputDir string) error {
	abs, err := filepath.Abs(outputDir)
	if err != nil {
		return err
	}
	grammarDir := filepath.Join(abs, "grammar")
	langs, err := listGrammarLangs(grammarDir)
	if err != nil {
		return err
	}

	if err := runGo(abs, "work", "sync"); err != nil {
		return fmt.Errorf("go work sync: %w", err)
	}
	if err := runGo(abs, "mod", "tidy"); err != nil {
		return fmt.Errorf("go mod tidy (root): %w", err)
	}
	if err := runGo(grammarDir, "mod", "tidy"); err != nil {
		return fmt.Errorf("go mod tidy (grammar): %w", err)
	}
	for _, lang := range langs {
		dir := filepath.Join(grammarDir, lang)
		if err := runGo(dir, "mod", "tidy"); err != nil {
			return fmt.Errorf("go mod tidy (grammar/%s): %w", lang, err)
		}
		slog.Info("tidied grammar module", "lang", lang)
	}
	return nil
}

func runGo(dir string, args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	return cmd.Run()
}

func listGrammarLangs(grammarDir string) ([]string, error) {
	entries, err := os.ReadDir(grammarDir)
	if err != nil {
		return nil, err
	}
	var langs []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		if _, err := os.Stat(filepath.Join(grammarDir, name, "api.go")); err != nil {
			continue
		}
		langs = append(langs, name)
	}
	sort.Strings(langs)
	return langs, nil
}

func writeCoreGoMod(grammarDir string) error {
	content := fmt.Sprintf(`module %s

go %s

require modernc.org/libc %s

replace modernc.org/libc => ../third-party/libc
`, grammarModulePath, moduleGoVersion, libcModuleVersion)
	return os.WriteFile(filepath.Join(grammarDir, "go.mod"), []byte(content), 0644)
}

func writeLangGoMod(grammarDir, lang string) error {
	content := fmt.Sprintf(`module %s/%s

go %s

require (
	%s %s
	modernc.org/libc %s
)

replace %s => ../

replace modernc.org/libc => ../../third-party/libc
`, grammarModulePath, lang, moduleGoVersion, grammarModulePath, localPseudoVer, libcModuleVersion, grammarModulePath)
	return os.WriteFile(filepath.Join(grammarDir, lang, "go.mod"), []byte(content), 0644)
}

func writeGoWork(outputDir string, langs []string) error {
	var b strings.Builder
	b.WriteString("go " + moduleGoVersion + "\n\n")
	b.WriteString("use (\n")
	b.WriteString("\t.\n")
	b.WriteString("\t./grammar\n")
	for _, lang := range langs {
		fmt.Fprintf(&b, "\t./grammar/%s\n", lang)
	}
	b.WriteString(")\n")
	return os.WriteFile(filepath.Join(outputDir, "go.work"), []byte(b.String()), 0644)
}

// updateRootGoMod adds require + replace for the core grammar module and each
// language module so cmd/parse can depend on them from the monorepo root.
func updateRootGoMod(outputDir string, langs []string) error {
	path := filepath.Join(outputDir, "go.mod")
	data, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		// Tests / empty trees: create a minimal root module.
		data = []byte("module " + rootModulePath + "\n\ngo " + moduleGoVersion + "\n")
	}
	f, err := modfile.Parse(path, data, nil)
	if err != nil {
		return err
	}

	// Collect module paths we own under grammar/.
	mods := []struct {
		path string
		dir  string
	}{
		{grammarModulePath, "./grammar"},
	}
	for _, lang := range langs {
		mods = append(mods, struct {
			path string
			dir  string
		}{grammarModulePath + "/" + lang, "./grammar/" + lang})
	}

	// Drop existing require/replace for our nested modules so we can re-add cleanly.
	owned := map[string]bool{}
	for _, m := range mods {
		owned[m.path] = true
	}
	var reqKeep []*modfile.Require
	for _, r := range f.Require {
		if owned[r.Mod.Path] {
			continue
		}
		reqKeep = append(reqKeep, r)
	}
	f.Require = reqKeep
	var repKeep []*modfile.Replace
	for _, r := range f.Replace {
		if owned[r.Old.Path] {
			continue
		}
		repKeep = append(repKeep, r)
	}
	f.Replace = repKeep

	for _, m := range mods {
		if err := f.AddRequire(m.path, localPseudoVer); err != nil {
			return err
		}
		if err := f.AddReplace(m.path, "", m.dir, ""); err != nil {
			return err
		}
	}
	// Ensure libc is required so the existing replace stays meaningful for tooling.
	if err := f.AddRequire("modernc.org/libc", libcModuleVersion); err != nil {
		return err
	}

	// Nested grammar modules are direct deps of root (cmd/parse blank-imports).
	for _, r := range f.Require {
		if owned[r.Mod.Path] || r.Mod.Path == "modernc.org/libc" {
			r.Indirect = false
		}
	}
	// Regroup direct vs indirect require blocks.
	f.SetRequire(f.Require)

	f.Cleanup()
	out, err := f.Format()
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0644)
}
