#!/usr/bin/env bash

set -e

VERSION="${1:-dev}"
OUTPUT="dist"

rm -rf "$OUTPUT"
mkdir -p "$OUTPUT"

build() {
    GOOS="$1" GOARCH="$2" \
    go build \
    -ldflags "-X main.version=$VERSION" \
    -o "$OUTPUT/$3" \
    ./cmd/url-open
}

build linux   amd64 url-open-linux-amd64
build linux   arm64 url-open-linux-arm64
build freebsd amd64 url-open-freebsd-amd64
build freebsd arm64 url-open-freebsd-arm64
build darwin  amd64 url-open-darwin-amd64
build darwin  arm64 url-open-darwin-arm64
build windows amd64 url-open-windows-amd64.exe
build windows arm64 url-open-windows-arm64.exe

(
    cd "$OUTPUT"
    sha256sum url-open-* > SHA256SUMS
)
