#!/bin/bash

# This will point to the location of install.sh wherever it is
# example /home/mew/dev/projects/steamctl/scripts
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
DIST_DIR="$ROOT_DIR/dist"

# Build variables
VERSION="dev"
ARCH="amd64"

# Build artifacts
LINUX_FILE_NAME="steamctl-linux-${ARCH}"
WINDOWS_FILE_NAME="steamctl-windows-${ARCH}.exe"
CHECKSUM_FILE_NAME="checksums.txt"
MANIFEST_FILE_NAME="manifest.txt"

LINUX_BIN="${DIST_DIR}/${LINUX_FILE_NAME}"
WINDOWS_BIN="${DIST_DIR}/${WINDOWS_FILE_NAME}"
CHECKSUMS_FILE="${DIST_DIR}/${CHECKSUM_FILE_NAME}"
BUILD_MANIFEST="${ROOT_DIR}/${MANIFEST_FILE_NAME}"

# cd to project directory
cd "$ROOT_DIR"

# Get version # using git tags
VERSION=$(git describe --tags --dirty --always)

# Override version only if arg is passed
if [[ $# -ge 1 ]]; then
  VERSION="$1"
fi

TIME=$(date -u -Is)
printf "  - timestamp: ${TIME}\n"
printf "    build:\n      version: ${VERSION}\n"

# Calculate a new hash of all git tracked files to check against manifest hash
BUILD_HASH=$(
  git ls-files -z \
  | grep -zv -E '(^scripts/|^dist/|_test\.go$|^README\.md$|^LICENSE$|^\.gitignore$|^\.env\.example$)' \
  | sort -z \
  | xargs -0 sha256sum \
  | sha256sum \
  | awk '{print $1}'
)

printf "      hash: ${BUILD_HASH}\n"

MANIFEST_HASH=""
if [[ -f "$BUILD_MANIFEST" ]]; then
  MANIFEST_HASH="$(cat "$BUILD_MANIFEST")"
fi

printf "      manifest: ${MANIFEST_HASH}\n"

if [[ "$BUILD_HASH" == "$MANIFEST_HASH" ]]; then
  printf "      diff: unchanged\n"
  printf "    pipeline_result: skip\n"
  exit 0
fi
printf "      diff: changed\n"

# Execute go tests
TEST_OUTPUT="$(go test ./... 2>&1)"
TEST_RESULT=$?
if [[ $TEST_RESULT -eq 0 ]]; then
  printf "    tests:\n"
  printf "      output: |\n"
  while IFS= read -r line; do
    printf "        %s\n" "$line"
  done <<< "$TEST_OUTPUT"
  printf "      result: pass\n"

  # Update manifest hash if changes detected and tests pass
  printf "%s\n" "$BUILD_HASH" > "$BUILD_MANIFEST"

  rm -f "$LINUX_BIN" "$WINDOWS_BIN" "$CHECKSUMS_FILE"

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
  
  # Checksums
  cd "$DIST_DIR"
  sha256sum steamctl-* > "$CHECKSUMS_FILE"

  want="$VERSION"
  got=$("$LINUX_BIN" --version)

  printf "    verify:\n"
  printf "      got: $got\n"
  printf "      want: $want\n"
  if [[ "$got" != *"$want"* ]]; then
    printf "      result: fail\n"
    printf "    pipeline_result: fail\n"
  else
    printf "      result: pass\n"
    
    ARTIFACTS=$(printf '%s ' *)
    printf "    artifacts:\n      got: ${ARTIFACTS}\n      want: ${CHECKSUM_FILE_NAME} ${LINUX_FILE_NAME} ${WINDOWS_FILE_NAME}\n"

    if [[ "$ARTIFACTS" == *"$LINUX_FILE_NAME"* && "$ARTIFACTS" == *"$WINDOWS_FILE_NAME"* && "$ARTIFACTS" == *"$CHECKSUM_FILE_NAME"* ]]; then
      printf "      result: pass\n" 
      printf "    pipeline_result: pass\n"
    else
    printf "      result: fail\n" 
    printf "    pipeline_result: fail\n"
    fi
  fi
else
  printf "    tests:\n"
  printf "      output: |\n"
  while IFS= read -r line; do
    printf "        %s\n" "$line"
  done <<< "$TEST_OUTPUT"
  printf "      result: fail\n"
  printf "    pipeline_result: fail\n"
fi