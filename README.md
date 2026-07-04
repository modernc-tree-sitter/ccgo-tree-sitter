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
| `codegen:darwin-arm64` | darwin/arm64 | `macos-latest` | `-fno-blocks`, empty nullability macros |
| `codegen:windows-amd64` | windows/amd64 | `windows-latest` | clang `--target=x86_64-w64-mingw32` + MinGW sysroot |
| `codegen:windows-arm64` | windows/arm64 | `windows-11-arm` | clang `--target=aarch64-w64-mingw32` (experimental) |

Outputs are `grammar/core-{GOOS}-{GOARCH}.go` and
`grammar/<lang>/grammar-{GOOS}-{GOARCH}.go` with matching `//go:build` tags.

Preprocessing uses **clang** (`CC=clang`). On Windows, codegen drives the **MinGW**
triple (not MSVC headers) and honors `MINGW_SYSROOT` when set. On Darwin it passes
`-fno-blocks` and defines `_Nonnull` / `_Nullable` / `_Null_unspecified` empty so
Apple qualifiers never reach ccgo. `CC` / `CFLAGS` are parsed with
[`shell.Fields`](https://pkg.go.dev/mvdan.cc/sh/v3/shell#Fields). Linux hosts still
drop clang/glibc `_Float*` typedefs after `-E`.
