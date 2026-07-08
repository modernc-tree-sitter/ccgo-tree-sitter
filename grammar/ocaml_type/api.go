package grammar_ocaml_type

import (
	"unsafe"
	"reflect"
	"github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar"
)

// Language returns the TSLanguage for ocaml_type with external scanner properly connected
func Language() *grammar.TSLanguage {
	ptr := tree_sitter_ocaml_type(nil)
	lang := (*grammar.TSLanguage)(unsafe.Pointer(ptr))

	// WORKAROUND: ccgo doesn't properly initialize function pointers in struct literals
	// Manually connect external scanner functions
	if lang.Fexternal_scanner.Fcreate == 0 {
		lang.Fexternal_scanner.Fcreate = reflect.ValueOf(tree_sitter_ocaml_type_external_scanner_create).Pointer()
		lang.Fexternal_scanner.Fdestroy = reflect.ValueOf(tree_sitter_ocaml_type_external_scanner_destroy).Pointer()
		lang.Fexternal_scanner.Fscan = reflect.ValueOf(tree_sitter_ocaml_type_external_scanner_scan).Pointer()
		lang.Fexternal_scanner.Fserialize = reflect.ValueOf(tree_sitter_ocaml_type_external_scanner_serialize).Pointer()
		lang.Fexternal_scanner.Fdeserialize = reflect.ValueOf(tree_sitter_ocaml_type_external_scanner_deserialize).Pointer()
	}

	return lang
}

func init() {
	grammar.Register("ocaml_type", Language())
}
