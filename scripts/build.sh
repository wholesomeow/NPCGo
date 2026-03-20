#!/bin/bash

set -Eeuo

#######################################
# HELPERS
#######################################
# Create Build directory if it doesn't exist
mkdir -p ./build

# Colours for fun and cool terminal output
GREEN='033[0;32m'
YELLOW='033[1;33m'
RED='033[0;31m'
NC='033[0m'

green() { echo -e "${GREEN}$1${NC}"; }
yellow() { echo -e "${YELLOW}$1${NC}"; }
red() { echo -e "${RED}$1${NC}"; }
info() { echo -e "$1"; }

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    red "Missing required command: $1"
    exit 1
  fi
}

trap 'red "Error on line $LINENO. Exiting."; exit 1' ERR

# Runs a set of prechecks to make sure everything that needs to be installed
# is actually installed properly.
prechecks() {
  require_cmd go
  require_cmd npm
  require_cmd wails

  green "Go: $(go version)"
  green "npm: $(npm --version)"
  green "Wails: $(wails version)"

  # This check is weak but it works
  if ! echo "$PATH" | grep -q "go/bin"; then
    yellow "Warning: go/bin not explicitly in PATH"
  fi

  # Wails environment validation
  wails doctor || {
    red "Wails doctor failed"
    exit 1
  }
}

#######################################
# CLI BUILD
#######################################
build_cli_main() {
  rm -f "$BUILD_DIR/npcgo"

  info "Running tests..."
  go test ./...

  info "Building CLI binary..."
  CGO_ENABLED=1 go build -o "$BUILD_DIR/npcgo" ./cmd

  green "Built: $BUILD_DIR/npcgo"
}

build_cli_devtools() {
  rm -f "$BUILD_DIR/devtools"

  info "Running tests..."
  go test ./...

  info "Building devtools..."
  CGO_ENABLED=1 go build -o "$BUILD_DIR/devtools" ./cmd/devtools

  green "Built: $BUILD_DIR/devtools"
}

#######################################
# WAILS
#######################################
wails_dev() {
  info "Starting Wails dev environment..."
  # Hot reload, frontend + backend
  wails dev
}

wails_build() {
  info "Building Wails app..."

  VERSION=$(git describe --tags --always || echo "dev")
  COMMIT=$(git rev-parse --short HEAD || echo "none")
  BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)

  wails build \
    -ldflags "-X main.version=$VERSION -X main.commit=$COMMIT -X main.buildTime=$BUILD_TIME"

  green "Wails build complete"
}

#######################################
# RELEASE
#######################################
release_cli() {
  rm -f "$BUILD_DIR/npcgo.exe" "$BUILD_DIR/npcgo-linux"

  info "Running tests..."
  go test ./...

  VERSION=$(git describe --tags --always || echo "dev")
  COMMIT=$(git rev-parse --short HEAD || echo "none")
  BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)

  LDFLAGS="-X main.version=$VERSION -X main.commit=$COMMIT -X main.buildTime=$BUILD_TIME"

  info "Building Linux binary..."
  GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="$LDFLAGS" -o "$BUILD_DIR/npcgo-linux" ./cmd
  green "Built: $BUILD_DIR/npcgo-linux"

  info "Building Windows binary..."
  GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="$LDFLAGS" -o "$BUILD_DIR/npcgo.exe" ./cmd
  green "Built: $BUILD_DIR/npcgo.exe"

  info "Compressing..."
  zip -j "$BUILD_DIR/npcgo-linux.zip" "$BUILD_DIR/npcgo-linux"
  zip -j "$BUILD_DIR/npcgo-windows.zip" "$BUILD_DIR/npcgo.exe"
}

release_wails() {
  info "Building Wails release binaries..."

  require_cmd zip

  VERSION=$(git describe --tags --always || echo "dev")
  COMMIT=$(git rev-parse --short HEAD || echo "none")
  BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)

  LDFLAGS="-X main.version=$VERSION -X main.commit=$COMMIT -X main.buildTime=$BUILD_TIME"

  # Clean previous builds
  rm -rf build/bin
  mkdir -p build/bin

  #######################################
  # Linux Build
  #######################################
  info "Building Wails Linux (amd64)..."
  GOOS=linux GOARCH=amd64 CGO_ENABLED=1 \
    wails build -platform linux/amd64 -ldflags "$LDFLAGS"

  # Move output (Wails dumps into ./build/bin)
  if [ -d "build/bin" ]; then
    zip -r build/wails-linux-amd64.zip build/bin/*
    green "Linux build packaged"
  else
    red "Linux build failed: no output directory"
    exit 1
  fi

  #######################################
  # Windows Build
  #######################################
  info "Building Wails Windows (amd64)..."
  GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
    wails build -platform windows/amd64 -ldflags "$LDFLAGS"

  if [ -d "build/bin" ]; then
    zip -r build/wails-windows-amd64.zip build/bin/*
    green "Windows build packaged"
  else
    red "Windows build failed: no output directory"
    exit 1
  fi

  green "Wails release builds complete"
}

#######################################
# Entry
#######################################
usage() {
  cat <<EOF
Usage: $0 <mode>

Modes:
  main-build       Build CLI tool
  dev-build        Build devtools CLI
  release-build    Cross-platform CLI builds
  wails-dev        Run Wails dev (hot reload)
  wails-build      Build Wails app
EOF
  exit 1
}

main() {
  [[ $# -eq 1 ]] || usage

  mkdir -p "$BUILD_DIR"

  prechecks

  case "$1" in
    main-build)     build_cli_main ;;
    dev-build)      build_cli_devtools ;;
    release-build)  release_cli ;;
    wails-dev)      wails_dev ;;
    wails-build)    wails_build ;;
    *)              usage ;;
  esac

  green "Done."
}

main "$@"