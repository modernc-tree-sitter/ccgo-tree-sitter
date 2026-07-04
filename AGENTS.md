# Guidelines
- Don't edit ./grammar directly. Those files are all code generated.
- Codegen preprocesses with clang by default; on Windows prefer MinGW gcc on PATH.
- `mise run codegen` is host-native only. Multi-platform regen is CI matrix + merge PR.
- Do not add MSVC-header regex sanitizers; use MinGW gcc -E and ccgo ignore flags.
