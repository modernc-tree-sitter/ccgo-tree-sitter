#!/usr/bin/env bash
# Run codegen for the host GOOS/GOARCH only.
set -euo pipefail
root="$(cd "$(dirname "$0")/.." && pwd)"
cd "$root"
goos="$(go env GOOS)"
goarch="$(go env GOARCH)"
task="codegen:${goos}-${goarch}"
if ! mise tasks ls 2>/dev/null | awk '{print $1}' | grep -qx "$task"; then
  echo "error: no mise task $task (host is ${goos}/${goarch})" >&2
  exit 1
fi
exec mise run "$task"
