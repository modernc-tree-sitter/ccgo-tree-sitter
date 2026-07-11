package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ProjectGenerator generates a complete Go project structure including a CLI utility.
// It scaffolds the necessary directories and entrypoints to consume the transpiled grammars.
type ProjectGenerator struct {
	// OutputDir is the destination directory for the generated project.
	OutputDir string
	// Grammars is a list of grammar names to include in the generated CLI support.
	Grammars []string
}

// Generate orchestrates the creation of the complete project structure,
// currently focusing on generating the CLI parsing utility.
func (pg *ProjectGenerator) Generate() error {

	// Create CLI
	if err := pg.generateCLI(); err != nil {
		return err
	}

	return nil
}

// generateCLI scaffolds a command-line tool in the target output directory
// that can parse files using any of the available transpiled grammars, automatically
// matching the grammar to the file extension.
func (pg *ProjectGenerator) generateCLI() error {
	cmdDir := filepath.Join(pg.OutputDir, "cmd", "parse")
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		return err
	}

	// Extract grammar names (clean for Go identifiers)
	grammarNames := make([]string, len(pg.Grammars))
	extensions := make([]string, len(pg.Grammars))
	for i, g := range pg.Grammars {
		name := extractGrammarName(g)
		grammarNames[i] = name
		// Map common extensions
		switch name {
		case "javascript":
			extensions[i] = "js"
		case "typescript":
			extensions[i] = "ts"
		case "python":
			extensions[i] = "py"
		default:
			extensions[i] = name
		}
	}

	content := pg.generateMainGo(grammarNames, extensions)
	return os.WriteFile(filepath.Join(cmdDir, "main.go"), []byte(content), 0644)
}

func (pg *ProjectGenerator) generateMainGo(grammarNames, extensions []string) string {
	var imports strings.Builder
	var caseStatements strings.Builder
	var helpText strings.Builder

	moduleName := filepath.Base(pg.OutputDir)
	if moduleName == "." {
		moduleName = "tree-sitter-go"
	}

	// Generate imports
	imports.WriteString("\t\"" + moduleName + "/core\"\n")
	for _, name := range grammarNames {
		imports.WriteString(fmt.Sprintf("\t%s \"%s/%s\"\n", name, moduleName, name))
	}

	// Generate case statements and help
	for i, name := range grammarNames {
		ext := extensions[i]
		caseStatements.WriteString(fmt.Sprintf("\tcase \"%s\":\n", ext))
		caseStatements.WriteString(fmt.Sprintf("\t\tlang = %s.Language()\n", name))
		helpText.WriteString(fmt.Sprintf("  - %s (.%s)\n", name, ext))
	}

	return fmt.Sprintf(`package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

%s)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	filename := os.Args[1]

	// Auto-detect language from extension
	ext := strings.TrimPrefix(filepath.Ext(filename), ".")
	var lang *core.TSLanguage

	switch ext {
%s	default:
		fmt.Fprintf(os.Stderr, "Unsupported file extension: .%%s\n", ext)
		fmt.Fprintf(os.Stderr, "Supported languages:\n%%s\n", supportedLanguages())
		os.Exit(1)
	}

	// Read file
	source, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %%v\n", err)
		os.Exit(1)
	}

	// Create parser (GC-managed; no Delete required)
	parser := core.NewParser()
	if !parser.SetLanguage(lang) {
		fmt.Fprintf(os.Stderr, "Failed to set language\n")
		os.Exit(1)
	}

	// Parse
	tree := parser.ParseString(string(source))

	// Print tree
	root := tree.RootNode()
	printNode(root, source, "", "")
}

func printNode(n *core.Node, source []byte, indent string, fieldName string) {
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

	fmt.Printf("%%s%%s%%s [%%d-%%d]", indent, prefix, typeStr, start, end)

	if n.ChildCount() == 0 {
		fmt.Printf(" %%q", string(source[start:end]))
	}
	fmt.Println()

	count := n.ChildCount()
	for i := uint32(0); i < count; i++ {
		child := n.Child(i)
		field := n.FieldNameForChild(i)
		printNode(child, source, indent+"  ", field)
	}
}

func printUsage() {
	fmt.Println("Usage: parse <file>")
	fmt.Println("\nSupported languages:")
	fmt.Print(supportedLanguages())
}

func supportedLanguages() string {
	return %s
}
`, imports.String(), caseStatements.String(), "`"+helpText.String()+"`")
}
