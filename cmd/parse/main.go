package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/lucasew/ccgo-tree-sitter/grammar"
	"github.com/lucasew/ccgo-tree-sitter/parsejson"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:          "parse <language> <file>",
	Short:        "Parse source files and print tree-sitter nodes",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	RunE: func(_ *cobra.Command, args []string) error {
		return run(args[0], args[1])
	},
}

func main() {
	Command.SetUsageTemplate(Command.UsageTemplate() + "\nSupported languages:\n" + supportedLanguages() + "\n")
	if err := Command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(languageName, filename string) error {
	// Read file
	source, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Resolve language
	lang, ok := grammar.Get(strings.ToLower(languageName))
	if !ok {
		return fmt.Errorf(
			"unsupported language: %s\nsupported languages: %s",
			languageName,
			supportedLanguages(),
		)
	}

	// Create parser
	parser := grammar.NewParser()
	defer parser.Delete()
	if !parser.SetLanguage(lang) {
		return fmt.Errorf("failed to set language")
	}

	// Parse
	tree := parser.ParseString(string(source))
	defer tree.Delete()

	// Convert parse tree to JSON
	root := tree.RootNode()
	output := parsejson.Output{
		Language: strings.ToLower(languageName),
		File:     filename,
		Root:     parsejson.BuildNode(root, source, ""),
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(output); err != nil {
		return fmt.Errorf("failed to encode output: %w", err)
	}
	return nil
}

func supportedLanguages() string {
	langs := grammar.List()
	if len(langs) == 0 {
		return "none"
	}
	sort.Strings(langs)
	return strings.Join(langs, ", ")
}
