# Guidelines
- Don't edit ./grammar directly. Those files are all code generated.
- Codegen preprocesses C with clang (`CC=clang` by default; override via `CC` / `CFLAGS`).
- Per-target tasks live in `mise.toml` as `codegen:<goos>-<goarch>`.
  `mise run codegen` is **host-native only** (`scripts/codegen-host.sh`).
  Full multi-platform regen is `.github/workflows/codegen.yml`: each runner
  transpiles its own triple, uploads an artifact, and `merge` opens the PR.
