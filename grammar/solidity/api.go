package grammar_solidity

import (
	"unsafe"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for solidity
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_solidity(nil)
	return (*grammar.TSLanguage)(unsafe.Pointer(ptr))
}

func init() {
	grammar.Register("solidity", Language())
}
