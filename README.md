# ccgo-tree-sitter
Tree sitter in Go using a patched modernc.org ccgo

Extremely experimental, but it can parse some stuff already.

## Codegen

Regenerate bindings for **this machine** (host `GOOS`/`GOARCH` only):

```bash
mise run codegen
```

Or one explicit triple:

```bash
mise run codegen:darwin-arm64
TARGET_GOOS=windows TARGET_GOARCH=amd64 go run ./cmd/codegen
```

`mise run codegen:all` runs every configured task on the current host; that only
works when headers/ABI match. Prefer the **CI matrix**, where each runner
transpiles its own native triple and a merge job opens a PR.

### Target matrix

| Task | `GOOS`/`GOARCH` | CI runner | Notes |
| --- | --- | --- | --- |
| `codegen:linux-amd64` | linux/amd64 | `ubuntu-latest` | checked in |
| `codegen:linux-arm64` | linux/arm64 | `ubuntu-24.04-arm` | checked in |
| `codegen:linux-386` | linux/386 (i386) | `ubuntu-latest` | experimental |
| `codegen:darwin-arm64` | darwin/arm64 | `macos-latest` | `-fno-blocks`, erase `__attribute__` / `API_AVAILABLE` |
| `codegen:windows-amd64` | windows/amd64 | `windows-latest` | MinGW triple + sysroot; erase `__attribute__` |
| `codegen:windows-arm64` | windows/arm64 | `windows-11-arm` | MinGW triple without x86_64 sysroot (experimental) |

Outputs are `grammar/core-{GOOS}-{GOARCH}.go` and
`grammar/<lang>/grammar-{GOOS}-{GOARCH}.go` with matching `//go:build` tags.

Preprocessing uses **clang** (`CC=clang`). On Windows it targets **MinGW**
(`*-w64-mingw32`) with an arch-matching `MINGW_SYSROOT` when available. On Darwin
and Windows, `__attribute__` / `__extension__` (and Apple `API_AVAILABLE`) are
defined empty so those tokens never reach ccgo—no source regex rewrites.
`CC` / `CFLAGS` use [`shell.Fields`](https://pkg.go.dev/mvdan.cc/sh/v3/shell#Fields).
Linux hosts still drop clang/glibc `_Float*` typedefs after `-E`.
