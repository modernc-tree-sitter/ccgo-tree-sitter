#!/bin/bash
set -ex

# Clone or update workspaced
if [ ! -d "/tmp/workspaced" ]; then
    git clone https://github.com/lucasew/workspaced.git /tmp/workspaced
else
    cd /tmp/workspaced
    git fetch
    git checkout origin/main
    cd -
fi

cd /tmp/workspaced

# Patch shellcheck.go to exclude third-party and grammar directories
find . -name "shellcheck.go" -exec sed -i -E 's/Exclude: \[\]string\{/Exclude: \[\]string\{"third-party", "grammar", /g' {} +
# Patch prelude.go to disable govulncheck and golangci to prevent OOM
find . -name "prelude.go" -exec sed -i -E '/golangci/d' {} +
find . -name "prelude.go" -exec sed -i -E '/govulncheck/d' {} +

# Install workspaced
go install ./cmd/workspaced
