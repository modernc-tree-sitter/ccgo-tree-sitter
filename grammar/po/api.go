package grammar_po

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for po
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_po(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("po", Language())
}
