module github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/psv

go 1.25.0

require (
	github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar v0.0.0-20260907205036-8fdfa3f25c9a
	modernc.org/libc v1.81.0
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/exp v0.0.0-20251023183803-a4bb9ffd2546 // indirect
	golang.org/x/sys v0.37.0 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
)

replace github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar => ../

replace modernc.org/libc => github.com/modernc-tree-sitter/libc v0.0.0-20260707203921-3c7a53d19f3f
