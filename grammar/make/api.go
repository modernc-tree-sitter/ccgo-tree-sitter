package grammar_make

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for make
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_make(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("make", Language())
}
