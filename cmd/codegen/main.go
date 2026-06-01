package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var (
	TREE_SITTER_PATH = "./third-party/tree-sitter"
	OUTPUT_DIR       = "."
	targetGOOS       string
	targetGOARCH     string
	keepTemp         bool
)

var rootCmd = &cobra.Command{
	Use:   "tree-sitter-go",
	Short: "Transpile tree-sitter C code to Go using ccgo",
	Long: `A tool to transpile tree-sitter core library and grammars from C to Go.

This tool uses ccgo to convert tree-sitter's C implementation into Go code,
allowing you to use tree-sitter parsers natively in Go without CGO.`,
	RunE: run,
}

func env(key, defaultValue string) string {
	if ret := os.Getenv(key); ret != "" {
		return ret
	}
	return defaultValue
}

func init() {
	rootCmd.Flags().StringVar(&targetGOOS, "goos", env("TARGET_GOOS", runtime.GOOS), "Target GOOS for generated code")
	rootCmd.Flags().StringVar(&targetGOARCH, "goarch", env("TARGET_GOARCH", runtime.GOARCH), "Target GOARCH for generated code")
	rootCmd.Flags().BoolVarP(&keepTemp, "keep-temp", "k", false, "Keep temporary files for debugging")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	slog.Info("compiling for target", "GOOS", targetGOOS, "GOARCH", targetGOARCH)
	// Create transpiler
	transpiler := &Transpiler{
		TreeSitterPath: TREE_SITTER_PATH,
		GOOS:           targetGOOS,
		GOARCH:         targetGOARCH,
		KeepTemp:       keepTemp,
	}

	// Transpile core
	slog.Info("transpiling tree-sitter core", "path", TREE_SITTER_PATH)
	coreOutput := filepath.Join(OUTPUT_DIR, "grammar")
	if err := transpiler.TranspileCore(coreOutput); err != nil {
		return fmt.Errorf("failed to transpile core: %w", err)
	}

	grammars, err := filepath.Glob("third-party/tree-sitter-*")
	if err != nil {
		return err
	}

	var summaryWriter io.Writer

	if summaryPath := os.Getenv("GITHUB_STEP_SUMMARY"); summaryPath != "" {
		summaryFile, err := os.OpenFile(summaryPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			slog.Warn("failed to open GITHUB_STEP_SUMMARY", "path", summaryPath, "error", err)
		} else {
			summaryWriter = summaryFile
			defer summaryFile.Close()
		}
	} else {
		var buf bytes.Buffer
		defer fmt.Println(buf.String())
		summaryWriter = &buf
	}

	fmt.Fprintf(summaryWriter, "## Grammar Codegen Summary\n\n")
	// Transpile grammars
	for i, grammarPath := range grammars {
		grammarName := extractGrammarName(grammarPath)
		slog.Info("transpiling grammar", "index", i+1, "total", len(grammars), "grammar", grammarName, "path", grammarPath)
		if err := transpiler.TranspileGrammar(grammarPath, OUTPUT_DIR+"/grammar"); err != nil {
			slog.Warn("failed to transpile grammar, skipping", "grammar", grammarName, "path", grammarPath, "error", err)
			fmt.Fprintf(summaryWriter, "- `%s/%s` `%s`: ❌\n", targetGOOS, targetGOARCH, grammarName)
			continue
		}
		fmt.Fprintf(summaryWriter, "- `%s/%s` `%s`: ✅\n", targetGOOS, targetGOARCH, grammarName)
	}

	slog.Info("updating languages registry in cmd/parse/languages.go")
	if err := updateLanguagesGo(OUTPUT_DIR); err != nil {
		return fmt.Errorf("failed to update languages registry: %w", err)
	}

	return nil
}

func updateLanguagesGo(outputDir string) error {
	grammarDir := filepath.Join(outputDir, "grammar")
	entries, err := os.ReadDir(grammarDir)
	if err != nil {
		return err
	}

	var languages []string
	for _, entry := range entries {
		if entry.IsDir() {
			languages = append(languages, entry.Name())
		}
	}
	sort.Strings(languages)

	moduleName := "github.com/lucasew/ccgo-tree-sitter"

	var sb strings.Builder
	sb.WriteString("package main\n\n")
	sb.WriteString("import (\n")
	for _, lang := range languages {
		fmt.Fprintf(&sb, "\t_ \"%s/grammar/%s\"\n", moduleName, lang)
	}
	sb.WriteString(")\n")

	targetFile := filepath.Join(outputDir, "cmd", "parse", "languages.go")
	return os.WriteFile(targetFile, []byte(sb.String()), 0644)
}
