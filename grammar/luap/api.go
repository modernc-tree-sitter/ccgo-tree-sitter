package grammar_luap

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for luap
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_luap(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("luap", Language())
}
