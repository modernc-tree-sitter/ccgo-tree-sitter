package grammar_puppet

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for puppet
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_puppet(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("puppet", Language())
}
