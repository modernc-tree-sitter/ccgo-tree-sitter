#!/bin/bash
set -euo pipefail

WORK_DIR=$(mktemp -d)
trap "rm -rf $WORK_DIR" EXIT

echo "Cloning workspaced..."
git clone https://github.com/lucasew/workspaced "$WORK_DIR"
cd "$WORK_DIR"
git checkout 08a6dce4000a

echo "Patching workspaced..."
# Patch shellcheck provider to ignore grammar and third-party directories to avoid OOM
sed -i 's/func(path string, d os.DirEntry, err error) error {/func(path string, d os.DirEntry, err error) error { if d.IsDir() \&\& (d.Name() == "grammar" || d.Name() == "third-party") { return filepath.SkipDir };/' pkg/provider/lint/shellcheck/shellcheck.go

# Disable govulncheck and golangci-lint by commenting out their registration in prelude.go
# These tools are failing due to Go version mismatch (project requires 1.25.0) or timeouts/OOM.
echo "Disabling problematic linters (govulncheck, golangci)..."
sed -i 's|_ "workspaced/pkg/provider/lint/govulncheck"|// _ "workspaced/pkg/provider/lint/govulncheck"|' pkg/provider/prelude/prelude.go
sed -i 's|_ "workspaced/pkg/provider/lint/golangci"|// _ "workspaced/pkg/provider/lint/golangci"|' pkg/provider/prelude/prelude.go

echo "Installing workspaced..."
go install ./cmd/workspaced
