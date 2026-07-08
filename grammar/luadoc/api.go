package grammar_luadoc

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for luadoc
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_luadoc(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("luadoc", Language())
}
