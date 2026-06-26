package grammar_glsl

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for glsl
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_glsl(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("glsl", Language())
}
