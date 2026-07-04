# ccgo-tree-sitter
Tree sitter in Go using a patched modernc.org ccgo

Extremely experimental, but it can parse some stuff already.

## Codegen

```bash
mise run codegen                  # host GOOS/GOARCH only
mise run codegen:darwin-arm64     # explicit triple
```

Prefer the **CI matrix** (each runner owns its native triple; merge opens a PR).

### Target matrix

| Task | `GOOS`/`GOARCH` | CI runner | Notes |
| --- | --- | --- | --- |
| `codegen:linux-amd64` | linux/amd64 | `ubuntu-latest` | checked in |
| `codegen:linux-arm64` | linux/arm64 | `ubuntu-24.04-arm` | checked in |
| `codegen:linux-386` | linux/386 | `ubuntu-latest` | experimental |
| `codegen:darwin-arm64` | darwin/arm64 | `macos-latest` | erase availability attrs; ccgo ignores align-16 |
| `codegen:windows-amd64` | windows/amd64 | `windows-latest` | **MinGW `gcc -E`** (not clang/MSVC) |

Outputs: `grammar/core-{GOOS}-{GOARCH}.go`, `grammar/<lang>/grammar-{GOOS}-{GOARCH}.go`.

On Windows, codegen prefers `x86_64-w64-mingw32-gcc` / `gcc` from MinGW on `PATH` even if
`CC=clang`, and passes `-ignore-unsupported-alignment` (and friends) to ccgo. On Darwin,
availability macros are defined empty and the same ccgo ignores apply. Linux still drops
`_Float*` typedefs after `-E`.
