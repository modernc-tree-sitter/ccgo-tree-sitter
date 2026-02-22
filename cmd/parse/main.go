package main

import (
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
	fmt.Printf("Using grammar: %p\n", lang)

	// Create parser
	parser := grammar.NewParser()
	defer parser.Delete()
	if !parser.SetLanguage(lang) {
		return fmt.Errorf("failed to set language")
	}

	// Parse
	tree := parser.ParseString(string(source))
	defer tree.Delete()

	// Print tree
	root := tree.RootNode()
	printNode(root, source, "", "")
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

func printNode(n *grammar.Node, source []byte, indent string, fieldName string) {
	if n.IsNull() {
		return
	}

	typeStr := n.Type()
	start := n.StartByte()
	end := n.EndByte()

	prefix := ""
	if fieldName != "" {
		prefix = fieldName + ": "
	}

	fmt.Printf("%s%s%s [%d-%d]", indent, prefix, typeStr, start, end)

	if n.ChildCount() == 0 {
		fmt.Printf(" %q", string(source[start:end]))
	}
	fmt.Println()

	count := n.ChildCount()
	for i := uint32(0); i < count; i++ {
		child := n.Child(i)
		field := n.FieldNameForChild(i)
		printNode(child, source, indent+"  ", field)
	}
}
