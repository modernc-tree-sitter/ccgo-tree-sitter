// This go.mod exists only to prevent the Go toolchain from treating
// any subdirectory under third-party/ as belonging to the parent module
// (github.com/lucasew/ccgo-tree-sitter).
//
// Most tree-sitter-* submodules have their own go.mod, which already
// makes them separate modules. A few (e.g. tensorflow, django) do not,
// and their .go files would otherwise be discovered by `go test ./...`,
// `go list ./...`, etc., causing build failures due to missing dependencies.
//
// The replaces in the root go.mod that point into third-party/ccgo/*
// continue to work because they are explicit path-based replaces.
module ignore

go 1.25
