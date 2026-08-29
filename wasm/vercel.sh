#!/usr/bin/env bash
#
# Builds the playground's WebAssembly binary. Run from this directory.
set -euo pipefail

GO_VERSION=1.24.1

dnf install -y wget tar gzip
wget -q "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz"
tar -C /usr/local -xzf "go${GO_VERSION}.linux-amd64.tar.gz"
export PATH=$PATH:/usr/local/go/bin

# GOOS and GOARCH have to be set here rather than left to the environment. The
# command without them does not merely produce a native binary, it fails to
# compile: -tags wasm pulls the wasm runtime files in beside the host ones and
# the two redeclare each other.
#
# Building the package rather than main.go alone keeps working if package main
# ever gains a second file.
GOOS=js GOARCH=wasm go build -tags wasm -o main.wasm ..
