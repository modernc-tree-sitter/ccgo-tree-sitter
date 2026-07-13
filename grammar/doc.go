// Package grammar binds tree-sitter for pure Go (modernc.org/ccgo, no CGO).
//
// # Ownership
//
// Parser, Tree, Query, and QueryCursor free native state via runtime cleanup
// when the Go value becomes unreachable. Delete is optional (eager free).
// A Tree pins its *Parser; a QueryCursor pins its *Query. Keep the parent
// reachable while using Nodes or other values derived from it.
//
// # Concurrency
//
// Language values are immutable after load and may be shared across parsers
// and goroutines. Parser and Query methods are safe for concurrent use: each
// serializes access to its native handle and TLS with an internal mutex.
// Callers need no external locking. For high parallel throughput, prefer one
// Parser or Query per goroutine to avoid contention on that mutex.
//
// # Registry
//
// Blank-import a language package (for example
// _ "github.com/modernc-tree-sitter/ccgo-tree-sitter/grammar/go") so its init
// calls Register. Look up languages with Get (by name) or GetByExtension.
// Call Register only for languages you load outside those packages.
//
// # Generated files
//
// core-*.go and per-language grammar-*.go are codegen output
// (mise run codegen). Do not edit them by hand; change sources or the
// generator and regenerate.
//
// # Line index
//
// LineIndex converts tree-sitter byte offsets to 1-based lines and 0-based
// columns (byte offset within the line). Build once per source with
// NewLineIndex or NewLineIndexBytes; lookups are O(log lines).
package grammar
