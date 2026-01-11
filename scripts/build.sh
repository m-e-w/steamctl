#!/bin/bash

DIR="${HOME}/dev/projects/steamctl/"
VERSION="dev"
OS="linux"
ARCH="amd64"
BIN="steamctl-${OS}-${ARCH}"

cd "$DIR"&&
VERSION=$(git describe --tags --dirty --always)

# Override version only if arg is passed
if [[ $# -ge 1 ]]; then
  VERSION="$1"
fi

go build -ldflags "-X github.com/m-e-w/steamctl/cmd.version=${VERSION}" -o "$BIN" &&
sha256sum "$BIN" > "checksums.txt"