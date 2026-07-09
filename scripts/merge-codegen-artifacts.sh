#!/usr/bin/env bash
# Merge per-target artifact dirs under $1 into the repo working tree.
set -euo pipefail
artifacts_root="${1:?usage: merge-codegen-artifacts.sh <artifacts-root>}"
root="$(cd "$(dirname "$0")/.." && pwd)"
cd "$root"

shopt -s nullglob
dirs=("$artifacts_root"/*/)
if [[ ${#dirs[@]} -eq 0 ]]; then
  echo "error: no artifacts in $artifacts_root" >&2
  exit 1
fi

preferred_dir=""
preferred_name=""
merged=0
for d in "${dirs[@]}"; do
  name="$(basename "${d%/}")"
  if [[ ! -d "$d/grammar" ]]; then
    echo "skipping $name (no grammar/)"
    continue
  fi
  echo "merging $name"
  cp -a "$d/grammar/." grammar/
  merged=$((merged + 1))
  # Prefer linux-amd64 for shared files (languages.go, go.work, root go.mod).
  if [[ "$name" == "linux-amd64" || "$name" == "codegen-linux-amd64" || -z "$preferred_dir" ]]; then
    preferred_dir="$d"
    preferred_name="$name"
  fi
done

if [[ "$merged" -eq 0 ]]; then
  echo "error: no artifact bundles contained grammar/" >&2
  exit 1
fi

if [[ -n "$preferred_dir" ]]; then
  if [[ -f "$preferred_dir/cmd/parse/languages.go" ]]; then
    mkdir -p cmd/parse
    cp "$preferred_dir/cmd/parse/languages.go" cmd/parse/languages.go
    echo "installed cmd/parse/languages.go from ${preferred_name}"
  fi
  # New grammars rewrite languages.go and need matching go.work / root go.mod.
  if [[ -f "$preferred_dir/go.work" ]]; then
    cp "$preferred_dir/go.work" go.work
    echo "installed go.work from ${preferred_name}"
  fi
  if [[ -f "$preferred_dir/go.mod" ]]; then
    cp "$preferred_dir/go.mod" go.mod
    echo "installed go.mod from ${preferred_name}"
  fi
  if [[ -f "$preferred_dir/go.sum" ]]; then
    cp "$preferred_dir/go.sum" go.sum
    echo "installed go.sum from ${preferred_name}"
  fi
fi

echo "merged ${merged} artifact bundles"
