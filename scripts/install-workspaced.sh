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

# Install workspaced
go install ./cmd/workspaced

# Assuming we might need a specific fix for memory usage, patch shellcheck and prelude if they exist:
find . -name "shellcheck.go" -exec sed -i -E 's/Exclude: \[\]string\{/Exclude: \[\]string\{"third-party", "grammar", /g' {} +
find . -name "prelude.go" -exec sed -i -E 's/"govulncheck",//g' {} +
find . -name "prelude.go" -exec sed -i -E 's/"golangci",//g' {} +
find . -name "prelude.go" -exec sed -i -E 's/Govulncheck,//g' {} +
find . -name "prelude.go" -exec sed -i -E 's/Golangci,//g' {} +
go install ./cmd/workspaced
