# ccgo-tree-sitter
Tree sitter in Go using a patched modernc.org ccgo

Extremely experimental, but it can parse some stuff already.

## Codegen

Regenerate checked-in `grammar/` sources with:

```bash
mise run codegen
```

That runs every `codegen:<goos>-<goarch>` task. One target only:

```bash
mise run codegen:darwin-arm64
# or
TARGET_GOOS=windows TARGET_GOARCH=amd64 go run ./cmd/codegen
```

### Target matrix

| Task | `GOOS`/`GOARCH` | Checked in? | Status |
| --- | --- | --- | --- |
| `codegen:linux-amd64` | linux/amd64 | yes | working |
| `codegen:linux-arm64` | linux/arm64 | yes | working |
| `codegen:linux-386` | linux/386 (i386) | no | mise wired; ccgo emits `&(*T)(unsafe.Pointer(...))` that `gc` rejects on 32-bit |
| `codegen:darwin-amd64` | darwin/amd64 | no | mise wired; blocked on target-correct preprocess (host glibc pulls `__ctype_b_loc` etc.) |
| `codegen:darwin-arm64` | darwin/arm64 | no | same as darwin/amd64 |
| `codegen:windows-amd64` | windows/amd64 | no | mise wired; blocked on target-correct preprocess (`windows.h` / no glibc shims in libc) |
| `codegen:windows-arm64` | windows/arm64 | no | same as windows/amd64 |

Outputs are `grammar/core-{GOOS}-{GOARCH}.go` and
`grammar/<lang>/grammar-{GOOS}-{GOARCH}.go` with matching `//go:build` tags.
Consumers select them automatically from `GOOS`/`GOARCH` (`CGO_ENABLED=0`).
Until a target is regenerated, that `GOOS`/`GOARCH` simply has no sources.

Next unblockers: preprocess against `modernc.org/libc` headers for the target
(`-nostdinc` + ccgo’s isystem) or run codegen on native/sysrooted runners; fix
32-bit address-of-pointer casts in generated code or patched ccgo.

Preprocessing uses **clang** (`CC=clang`). Override with `CC` (e.g. `zig cc`) and
`CFLAGS`. Both are parsed with
[`shell.Fields`](https://pkg.go.dev/mvdan.cc/sh/v3/shell#Fields) (POSIX words,
quotes, `$VAR`). After `-E`, codegen strips clang/glibc `_Float*` / `__float128`
typedefs that ccgo rejects.

By default tasks do **not** pass `--target`; host headers are preprocessed and
ccgo applies the target ABI via `modernc.org/libc`. For stricter triples, set
e.g. `CFLAGS=--target=x86_64-apple-darwin` **with** a matching sysroot
(`-isysroot` / `-sysroot`). Runtime support requires a `modernc.org/libc` port
for that pair.
