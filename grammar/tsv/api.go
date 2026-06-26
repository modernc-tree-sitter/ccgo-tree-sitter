package grammar_tsv

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for tsv
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_tsv(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("tsv", Language())
}
