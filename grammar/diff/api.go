package grammar_diff

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for diff
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_diff(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("diff", Language())
}
