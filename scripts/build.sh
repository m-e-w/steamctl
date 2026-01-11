#!/bin/bash

set -e 

# This will point to the location of install.sh wherever it is
# example /home/mew/dev/projects/steamctl/scripts
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
DIST_DIR="$ROOT_DIR/dist"

# Build variables
VERSION="dev"
ARCH="amd64"

# Build artifacts
LINUX_BIN="${DIST_DIR}/steamctl-linux-${ARCH}"
WINDOWS_BIN="${DIST_DIR}/steamctl-windows-${ARCH}.exe"
CHECKSUMS_FILE="${DIST_DIR}/checksums.txt"

# cd to project directory
cd "$ROOT_DIR"

# Get version # using git tags
VERSION=$(git describe --tags --dirty --always)

# Override version only if arg is passed
if [[ $# -ge 1 ]]; then
  VERSION="$1"
fi

echo "Build Version: $VERSION"

# Execute go tests
echo "Running go tests"

go test ./...

echo "All tests passed. Generating build artifacts"

echo "Cleaning up old build artifacts"
rm -f "$LINUX_BIN" "$WINDOWS_BIN" "$CHECKSUMS_FILE"

# Linux build
echo "Building linux binary"
GOOS=linux GOARCH=$ARCH \
go build \
  -ldflags "-X github.com/m-e-w/steamctl/cmd.version=${VERSION}" \
  -o "$LINUX_BIN"

# Windows build
echo "Building windows binary"
GOOS=windows GOARCH=$ARCH \
go build \
  -ldflags "-X github.com/m-e-w/steamctl/cmd.version=${VERSION}" \
  -o "$WINDOWS_BIN"

# Checksums
echo "Generating checksums.txt file"
cd "$DIST_DIR"
sha256sum steamctl-* > "$CHECKSUMS_FILE"