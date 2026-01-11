#!/bin/bash

STEAMCTL_ROOT_DIR="${HOME}/dev/projects/steamctl/"
STEAMCTL_BUILD_VERSION="dev"

cd "$STEAMCTL_ROOT_DIR"&&
STEAMCTL_BUILD_VERSION=$(git describe --tags --dirty --always)

# Override version only if arg is passed
if [[ $# -ge 1 ]]; then
  STEAMCTL_BUILD_VERSION="$1"
fi

go build -ldflags "-X github.com/m-e-w/steamctl/cmd.version=${STEAMCTL_BUILD_VERSION}" -o "steamctl-linux-amd64"