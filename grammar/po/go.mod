module github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/po

go 1.25.0

require (
	github.com/modernc-tree-sitter/ccgo-tree-sitter/core v0.0.0-20260907205036-8fdfa3f25c9a
	modernc.org/libc v1.67.6
)

replace github.com/modernc-tree-sitter/ccgo-tree-sitter/core => ../../core

replace modernc.org/libc => github.com/modernc-tree-sitter/libc v0.0.0-20260707203921-3c7a53d19f3f
