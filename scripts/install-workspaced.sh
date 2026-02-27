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
# We inject the skip logic before the error check in the WalkDir callback.
# This affects both Detect and Run methods in pkg/provider/lint/shellcheck/shellcheck.go

# Replace only inside the filepath.WalkDir callback function (lines 41 and 76 in original file)
# The pattern `if err != nil {` appears multiple times, so we need to be more specific or use line numbers if possible,
# but sed line numbers on a moving target (git repo) is risky.
# Instead, we will replace `if err != nil {` with the check ONLY when it is inside the callback.
# The callback signature is `func(path string, d os.DirEntry, err error) error {`

# We will use a slightly more robust sed command that matches the context of the callback start.

# Detect method callback
sed -i 's/func(path string, d os.DirEntry, err error) error {/func(path string, d os.DirEntry, err error) error { if d.IsDir() \&\& (d.Name() == "grammar" || d.Name() == "third-party") { return filepath.SkipDir };/' pkg/provider/lint/shellcheck/shellcheck.go

echo "Installing workspaced..."
go install ./cmd/workspaced
