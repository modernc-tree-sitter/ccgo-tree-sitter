package grammar_psv

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for psv
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_psv(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("psv", Language())
}
