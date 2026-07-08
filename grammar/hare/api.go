package grammar_hare

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for hare
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_hare(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("hare", Language())
}
