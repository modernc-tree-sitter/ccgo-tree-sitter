package grammar_rego

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/core"
)

// Language returns the TSLanguage for rego
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_rego(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("rego", Language())
}
