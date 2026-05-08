#!/usr/bin/env bash
set -e

WORKSPACED_DIR=$(mktemp -d)
git clone https://github.com/lucasew/workspaced.git "$WORKSPACED_DIR"
cd "$WORKSPACED_DIR"
go install -v ./cmd/workspaced
