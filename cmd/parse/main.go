package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/lucasew/ccgo-tree-sitter/grammar"
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

var querySource string
var queryFile string

func init() {
	Command.Flags().StringVar(&querySource, "query", "", "Tree-sitter query source")
	Command.Flags().StringVar(&queryFile, "query-file", "", "Path to a file containing a tree-sitter query")
}

func main() {
	Command.SetUsageTemplate(Command.UsageTemplate() + "\nSupported languages:\n" + supportedLanguages() + "\n")
	if err := Command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(languageName, filename string) error {
	if querySource != "" && queryFile != "" {
		return fmt.Errorf("use either --query or --query-file, not both")
	}

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

	query := querySource
	if queryFile != "" {
		queryBytes, err := os.ReadFile(queryFile)
		if err != nil {
			return fmt.Errorf("error reading query file: %w", err)
		}
		query = string(queryBytes)
	}

	// Convert parse tree to JSON
	root := tree.RootNode()
	if query != "" {
		compiledQuery, err := grammar.NewQuery(lang, query)
		if err != nil {
			return err
		}
		defer compiledQuery.Delete()
		matches := compiledQuery.ExecuteMatches(root, source)
		enc := json.NewEncoder(os.Stdout)
		for _, match := range matches {
			if err := enc.Encode(match); err != nil {
				return fmt.Errorf("failed to encode match output: %w", err)
			}
		}
		return nil
	}

	output := grammar.ParseOutput{
		Language: strings.ToLower(languageName),
		File:     filename,
		Root:     grammar.BuildParseNode(root, source, ""),
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
