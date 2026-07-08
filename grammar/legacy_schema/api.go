package grammar_legacy_schema

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for legacy_schema
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_legacy_schema(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("legacy_schema", Language())
}
