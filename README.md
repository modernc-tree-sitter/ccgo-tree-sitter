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

| Task | `GOOS`/`GOARCH` | CI runner | Checked in? |
| --- | --- | --- | --- |
| `codegen:linux-amd64` | linux/amd64 | `ubuntu-latest` | yes |
| `codegen:linux-arm64` | linux/arm64 | `ubuntu-24.04-arm` | yes |
| `codegen:linux-386` | linux/386 (i386) | `ubuntu-latest` (experimental) | no |
| `codegen:darwin-amd64` | darwin/amd64 | `macos-13` | no |
| `codegen:darwin-arm64` | darwin/arm64 | `macos-latest` | no |
| `codegen:windows-amd64` | windows/amd64 | `windows-latest` | no |
| `codegen:windows-arm64` | windows/arm64 | `windows-11-arm` (experimental) | no |

Outputs are `grammar/core-{GOOS}-{GOARCH}.go` and
`grammar/<lang>/grammar-{GOOS}-{GOARCH}.go` with matching `//go:build` tags.
Consumers select them via `GOOS`/`GOARCH` (`CGO_ENABLED=0`).

Preprocessing uses **clang** (`CC=clang`) on the runner that owns the triple.
Override with `CC` / `CFLAGS` (parsed by
[`shell.Fields`](https://pkg.go.dev/mvdan.cc/sh/v3/shell#Fields)). After `-E`,
codegen strips clang/glibc `_Float*` / `__float128` typedefs that ccgo rejects.
