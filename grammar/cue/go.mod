module github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/cue

go 1.25.0

require (
	github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar v0.0.0-20260907205036-8fdfa3f25c9a
	modernc.org/libc v1.67.6
)

replace github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar => ../

replace modernc.org/libc => github.com/modernc-tree-sitter/libc v1.73.5
