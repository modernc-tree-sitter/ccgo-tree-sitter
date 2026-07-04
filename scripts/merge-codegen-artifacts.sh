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

preferred_registry=""
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
  if [[ -f "$d/cmd/parse/languages.go" ]]; then
    if [[ "$name" == "linux-amd64" || -z "$preferred_registry" ]]; then
      preferred_registry="$d/cmd/parse/languages.go"
      preferred_name="$name"
    fi
  fi
done

if [[ "$merged" -eq 0 ]]; then
  echo "error: no artifact bundles contained grammar/" >&2
  exit 1
fi

if [[ -n "$preferred_registry" ]]; then
  mkdir -p cmd/parse
  cp "$preferred_registry" cmd/parse/languages.go
  echo "installed cmd/parse/languages.go from ${preferred_name}"
fi

echo "merged ${merged} artifact bundles"
