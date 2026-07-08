module github.com/lucasew/ccgo-tree-sitter/grammar/svelte

go 1.25.0

require (
	github.com/lucasew/ccgo-tree-sitter/grammar v0.0.0
	modernc.org/libc v1.67.6
)

replace github.com/lucasew/ccgo-tree-sitter/grammar => ../

replace modernc.org/libc => ../../third-party/libc
