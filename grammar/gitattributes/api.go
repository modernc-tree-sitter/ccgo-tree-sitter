package grammar_gitattributes

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for gitattributes
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_gitattributes(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("gitattributes", Language())
}
