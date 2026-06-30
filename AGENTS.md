# Tree-sitter Go Codegen

This repository transpiles tree-sitter C code to Go using ccgo.

## Error Handling
The project MUST have a single, centralized error-reporting function (e.g. `reportError`, `captureException`, or equivalent). All code paths that handle unexpected errors MUST funnel through this function — never call `console.error` or `os.Exit(1)` directly at the call site.
- The centralized error-reporting function is implemented as `reporter.HandleError` in `internal/reporter/reporter.go`.
- Cobra commands should use `SilenceErrors: true` to defer error logging to this centralized reporter.
- Ensure that `main()` functions delegate error handling to `reporter.HandleError` instead of using `os.Exit(1)` directly or logging raw errors.

## Execution Requirements
- The project uses Go 1.26.0 as the standardized version pinned via `mise.toml`.
- Use `workspaced` via `mise` for all linting and formatting needs. Do not manually install individual linters.
