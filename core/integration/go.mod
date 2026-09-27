module github.com/modernc-tree-sitter/ccgo-tree-sitter/core/integration

go 1.25.0

require (
	github.com/modernc-tree-sitter/ccgo-tree-sitter/core v0.0.0-20260927183436-febb7f868b50
	github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/go v0.0.0-20260927183534-cc64727aba48
	github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/json v0.0.0-20260927183534-cc64727aba48
	modernc.org/libc v1.67.6
)

replace github.com/modernc-tree-sitter/ccgo-tree-sitter/core => ../

replace modernc.org/libc => github.com/modernc-tree-sitter/libc v0.0.0-20260707203921-3c7a53d19f3f
