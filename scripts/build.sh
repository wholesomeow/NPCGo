#!/bin/bash

set -e

if [ -z $1 ]; then
  echo "Must provide build mode as argument"
  echo "Options are:"
  echo "  - main-build          Builds the main binary and runs its tests"
  echo "  - dev-build           Builds the devtools binary and runs its tests"
  echo "  - release-build       Builds the main binary for windows and linux"
  exit 1
fi

# Create Build directory if it doesn't exist
mkdir -p ./build

# Clean previous builds
rm -f build/npcgen build/npcgen.exe build/npcgen-linux build/devtools

case $1 in
  main-build )
    # Run the tests first so the binary wont build if tests fail
    echo "Running NPC Generator tests"
    go test ./...

    echo "Building NPC Generator Binary"
    CGO_ENABLED=0 go build -o build/npcgen ./cmd
    echo "Binary built at: build/npcgen-linux"
    ;;
  dev-build )
    # Run the tests first so the binary wont build if tests fail
    echo "Running NPC Generator Devtools tests"
    go test ./...

    echo "Building NPC Generator Binary"
    CGO_ENABLED=0 go build -o build/devtools ./cmd/devtools
    echo "Binary built at: build/devtools"
    ;;
  release-build )
    # Run the tests first so the binary wont build if tests fail
    echo "Running NPC Generator tests"
    go test ./...

    echo "Building NPC Generator Binary"
    # Linux 64-bit
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/npcgen-linux
    echo "Binary built at: build/npcgen-linux"

    # Windows 64-bit
    GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o build/npcgen.exe
    echo "Binary built at: build/npcgen.exe"

    echo "Compressing builds..."
    zip build/npcgen-linux.zip build/npcgen-linux
    zip build/npcgen-windows.zip build/npcgen.exe
    ;;
  
esac

echo "Build finished"