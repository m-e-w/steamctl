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

printf "Build Version: $VERSION\n\n"

# Execute go tests
printf "Running go tests\n"

go test ./...
printf "All tests passed\n\n"

printf "Cleaning up old build artifacts\n\n"
rm -f "$LINUX_BIN" "$WINDOWS_BIN" "$CHECKSUMS_FILE"

printf "Generating new build artifacts\n"
# Linux build
printf "Building linux binary\n"
GOOS=linux GOARCH=$ARCH \
go build \
  -ldflags "-X github.com/m-e-w/steamctl/cmd.version=${VERSION}" \
  -o "$LINUX_BIN"

# Windows build
printf "Building windows binary\n"
GOOS=windows GOARCH=$ARCH \
go build \
  -ldflags "-X github.com/m-e-w/steamctl/cmd.version=${VERSION}" \
  -o "$WINDOWS_BIN"

# Checksums
printf "Generating checksums.txt file\n"
cd "$DIST_DIR"
sha256sum steamctl-* > "$CHECKSUMS_FILE"


printf "\nTesting linux binary\n"

want="$VERSION"
got=$("$LINUX_BIN" --version)

printf "got:    $got\n"
printf "expect: $want\n"

if [[ "$got" != *"$want"* ]]; then
  printf "ERROR: version mismatch\n"
  printf "\nbuild failed\n"
  exit 1
fi

printf "\nBuild: Pass\n"