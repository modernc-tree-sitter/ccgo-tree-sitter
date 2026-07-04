# Guidelines
- Don't edit ./grammar directly. Those files are all code generated.
- Codegen preprocesses C with clang (`CC=clang` by default; override via `CC` / `CFLAGS`).
- Target matrix lives in `mise.toml` as `codegen:<goos>-<goarch>` tasks (`mise run codegen` runs all).
  Checked-in artifacts today may lag newly added tasks until the next full regen.
