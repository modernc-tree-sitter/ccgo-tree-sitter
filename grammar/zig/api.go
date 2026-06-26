package grammar_zig

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for zig
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_zig(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("zig", Language())
}
