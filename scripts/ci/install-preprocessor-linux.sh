#!/usr/bin/env bash
# Install clang (and multilib for linux/386) on Linux CI runners.
set -euo pipefail
goarch="${1:-$(uname -m)}"
sudo apt-get update
sudo apt-get install -y clang
case "$goarch" in
  386|i386|i686)
    sudo apt-get install -y gcc-multilib g++-multilib libc6-dev-i386
    ;;
esac
clang --version
