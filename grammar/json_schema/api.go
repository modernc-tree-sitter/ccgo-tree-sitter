package grammar_json_schema

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for json_schema
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_json_schema(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("json_schema", Language())
}
