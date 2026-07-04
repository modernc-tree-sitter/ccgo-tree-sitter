# ccgo-tree-sitter
Tree sitter in Go using a patched modernc.org ccgo

Extremely experimental, but it can parse some stuff already.

## Codegen

Regenerate checked-in `grammar/` sources with:

```bash
mise run codegen
```

Preprocessing uses **clang** by default (`CC=clang`). Override with `CC` (multi-word
launchers like `zig cc` are fine) and extra flags via `CFLAGS`. Both are parsed with
[`shell.Fields`](https://pkg.go.dev/mvdan.cc/sh/v3/shell#Fields) (POSIX word
splitting, quotes, and `$VAR` expansion), e.g.
`CFLAGS=--target=aarch64-unknown-linux-gnu`. After `-E`, codegen strips clang/glibc
`_Float*` / `__float128` typedefs that ccgo cannot type-check (GCC omits them as
builtins). Target selection for the generated Go is `TARGET_GOOS` / `TARGET_GOARCH`
or `--goos` / `--goarch`.
