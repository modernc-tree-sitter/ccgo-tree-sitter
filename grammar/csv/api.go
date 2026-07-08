package grammar_csv

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for csv
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_csv(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("csv", Language())
}
