#!/bin/bash

REPO="m-e-w/steamctl"
VERSION="latest"
OS="linux"
ARCH="amd64"
DIR="$HOME/.local/bin"
BIN="steamctl-${OS}-${ARCH}"
BASE_URL="https://github.com/${REPO}/releases/${VERSION}/download"

# Download binary and checksums
curl -fsLO "${BASE_URL}/${BIN}"
curl -fsLO "${BASE_URL}/checksums.txt"

# Verify checksum (Linux)
if grep " ${BIN}$" checksums.txt | sha256sum -c -; then
    : # Nothing to do
else
  exit 1
fi

# Check if user bin folder exists. Create it if it doesn't
if [[ -d "$DIR" ]]; then
    : # Nothing to do
else
    mkdir -p "$DIR"
fi

# Copy to users bin
cp "$BIN" "${DIR}/steamctl"

# Make it executable
chmod +x "${DIR}/steamctl"

# Clean up temp files
rm "$BIN" "checksums.txt"

echo "steamctl was installed to $DIR"