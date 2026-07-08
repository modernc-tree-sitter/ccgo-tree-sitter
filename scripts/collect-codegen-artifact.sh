#!/usr/bin/env bash
# Copy generated sources for one GOOS/GOARCH into $1 (artifact root).
set -euo pipefail
out="${1:?usage: collect-codegen-artifact.sh <outdir> <goos> <goarch>}"
goos="${2:?}"
goarch="${3:?}"
root="$(cd "$(dirname "$0")/.." && pwd)"
cd "$root"

rm -rf "$out"
mkdir -p "$out/grammar"

core="grammar/core-${goos}-${goarch}.go"
if [[ ! -f "$core" ]]; then
  echo "error: missing $core (codegen produced no core for ${goos}/${goarch})" >&2
  exit 1
fi
cp "$core" "$out/grammar/"

found=0
while IFS= read -r f; do
  [[ -z "$f" ]] && continue
  rel="${f#grammar/}"
  dir="$out/grammar/$(dirname "$rel")"
  mkdir -p "$dir"
  cp "$f" "$dir/"
  found=$((found + 1))
done < <(find grammar -type f -name "grammar-${goos}-${goarch}.go" | LC_ALL=C sort)

if [[ "$found" -eq 0 ]]; then
  echo "error: no grammar-*-${goos}-${goarch}.go files generated" >&2
  exit 1
fi

if [[ -f cmd/parse/languages.go ]]; then
  mkdir -p "$out/cmd/parse"
  cp cmd/parse/languages.go "$out/cmd/parse/"
fi
while IFS= read -r f; do
  [[ -z "$f" ]] && continue
  rel="${f#grammar/}"
  dir="$out/grammar/$(dirname "$rel")"
  mkdir -p "$dir"
  cp "$f" "$dir/"
done < <(find grammar -type f -name api.go | LC_ALL=C sort)

# Nested modules: core + per-lang go.mod (and go.work at repo root if present)
if [[ -f grammar/go.mod ]]; then
  cp grammar/go.mod "$out/grammar/"
fi
while IFS= read -r f; do
  [[ -z "$f" ]] && continue
  rel="${f#grammar/}"
  dir="$out/grammar/$(dirname "$rel")"
  mkdir -p "$dir"
  cp "$f" "$dir/"
done < <(find grammar -type f -name go.mod | LC_ALL=C sort)
if [[ -f go.work ]]; then
  cp go.work "$out/"
fi

echo "collected core + ${found} grammars for ${goos}/${goarch}"
