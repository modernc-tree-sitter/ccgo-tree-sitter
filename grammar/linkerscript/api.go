package grammar_linkerscript

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/core"
)

// Language returns the TSLanguage for linkerscript
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_linkerscript(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("linkerscript", Language())
}
