package main

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

var structPaddingPattern = regexp.MustCompile(`&struct \{ _ \[`)

type fixtureExpectation struct {
	rootType       string
	firstChildType string
}

var languageExpectations = map[string]fixtureExpectation{
	"go": {
		// Real-world fixture: refactree internal/fuzzy/catalog.go
		rootType:       "source_file",
		firstChildType: "package_clause",
	},
	"json": {
		rootType:       "document",
		firstChildType: "object",
	},
	"lua": {
		rootType: "chunk",
	},
}

func TestGeneratedCoreHasNoStructPadding(t *testing.T) {
	root := repoRoot(t)
	grammarDir := filepath.Join(root, "grammar")

	entries, err := os.ReadDir(grammarDir)
	if err != nil {
		t.Fatalf("failed to read grammar directory: %v", err)
	}

	var coreFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, "core-") && strings.HasSuffix(name, ".go") {
			coreFiles = append(coreFiles, filepath.Join(grammarDir, name))
		}
	}
	if len(coreFiles) == 0 {
		t.Fatalf("no generated core files found in %s", grammarDir)
	}

	for _, coreFile := range coreFiles {
		content, err := os.ReadFile(coreFile)
		if err != nil {
			t.Fatalf("failed to read %s: %v", coreFile, err)
		}
		if structPaddingPattern.Match(content) {
			t.Fatalf("found struct padding pattern in %s", coreFile)
		}
	}
}

func TestLanguageFixtures(t *testing.T) {
	root := repoRoot(t)
	fixturesRoot := filepath.Join(root, "testdata", "fixtures")

	langsWithFixtures := make(map[string]struct{})
	fixtureCount := 0
	err := filepath.WalkDir(fixturesRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == fixturesRoot {
			return nil
		}

		base := d.Name()
		if strings.HasPrefix(base, ".") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(fixturesRoot, path)
		if err != nil {
			return err
		}
		parts := strings.Split(filepath.ToSlash(relPath), "/")
		if len(parts) < 2 {
			return nil
		}

		language := parts[0]
		langsWithFixtures[language] = struct{}{}
		expectation, ok := languageExpectations[language]
		if !ok {
			return &testFailure{message: "missing expectation for fixture language", language: language}
		}

		lang, ok := grammar.Get(language)
		if !ok {
			return &testFailure{message: "language is not registered", language: language}
		}

		fixtureCount++
		fixturePath := path
		fixtureLabel := strings.Join(parts[1:], "/")
		t.Run(language+"/"+fixtureLabel, func(t *testing.T) {
			source, err := os.ReadFile(fixturePath)
			if err != nil {
				t.Fatalf("failed to read fixture %s: %v", fixturePath, err)
			}

			parser := grammar.NewParser()
			if !parser.SetLanguage(lang) {
				t.Fatalf("failed to set language %q", language)
			}

			tree := parser.ParseString(string(source))
			rootNode := tree.RootNode()
			if rootNode.IsNull() {
				t.Fatalf("root node is null for fixture %s", fixturePath)
			}
			if rootNode.HasError() {
				t.Fatalf("parse tree has errors for fixture %s:\n%s", fixturePath, rootNode.PrintTree())
			}

			parseNode := grammar.BuildParseNode(rootNode, source, "")
			if parseNode == nil {
				t.Fatalf("parse node is nil for fixture %s", fixturePath)
			}
			if parseNode.Type != expectation.rootType {
				t.Fatalf("unexpected root type for fixture %s: got %q want %q", fixturePath, parseNode.Type, expectation.rootType)
			}

			if expectation.firstChildType != "" {
				if len(parseNode.Children) == 0 {
					t.Fatalf("expected at least one child node for fixture %s", fixturePath)
				}
				if parseNode.Children[0].Type != expectation.firstChildType {
					t.Fatalf("unexpected first child type for fixture %s: got %q want %q", fixturePath, parseNode.Children[0].Type, expectation.firstChildType)
				}
			}
		})
		return nil
	})
	if err != nil {
		var testErr *testFailure
		if errors.As(err, &testErr) {
			if testErr.message == "language is not registered" {
				t.Fatalf("%v; supported languages: %s", err, supportedLanguages())
			}
			t.Fatal(err)
		}
		t.Fatalf("failed to discover fixtures: %v", err)
	}
	if fixtureCount == 0 {
		t.Fatalf("no fixtures found under %s", fixturesRoot)
	}

	for language := range languageExpectations {
		if _, ok := langsWithFixtures[language]; !ok {
			t.Fatalf("no fixtures found for language %q", language)
		}
	}
}

type testFailure struct {
	message  string
	language string
}

func (e *testFailure) Error() string {
	return e.message + `: "` + e.language + `"`
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatalf("could not locate repository root from %s", dir)
		}
		dir = parent
	}
}
