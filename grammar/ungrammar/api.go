package grammar_ungrammar

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/core"
)

// Language returns the TSLanguage for ungrammar
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_ungrammar(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("ungrammar", Language())
}
