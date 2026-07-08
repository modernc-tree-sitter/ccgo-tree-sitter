package grammar_core_schema

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for core_schema
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_core_schema(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("core_schema", Language())
}
