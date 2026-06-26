package grammar_commonlisp

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for commonlisp
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_commonlisp(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("commonlisp", Language())
}
