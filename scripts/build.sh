#!/bin/bash

DIR="${HOME}/dev/projects/steamctl/"
VERSION="dev"
ARCH="amd64"

OUT="dist"
LINUX_BIN="${OUT}/steamctl-linux-${ARCH}"
WINDOWS_BIN="${OUT}/steamctl-windows-${ARCH}.exe"

cd "$DIR"&&
VERSION=$(git describe --tags --dirty --always)

# Override version only if arg is passed
if [[ $# -ge 1 ]]; then
  VERSION="$1"
fi

# Linux build
GOOS=linux GOARCH=$ARCH \
go build \
  -ldflags "-X github.com/m-e-w/steamctl/cmd.version=${VERSION}" \
  -o "$LINUX_BIN"

# Windows build
GOOS=windows GOARCH=$ARCH \
go build \
  -ldflags "-X github.com/m-e-w/steamctl/cmd.version=${VERSION}" \
  -o "$WINDOWS_BIN"

cd "$OUT"&&
sha256sum steamctl-* > checksums.txt