package grammar_requirements

import (
	"unsafe"
	"github.com/lucasew/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for requirements
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_requirements(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("requirements", Language())
}
