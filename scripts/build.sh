#!/bin/bash

if [ -z $1 ]; then
  echo "Must provide build mode as argument"
  echo "Options are:"
  echo "  - main-build          Builds the main binary and runs its tests"
  echo "  - dev-build           Builds the devtools binary and runs its tests"
  echo "  - release-build       Builds the main binary for windows and linux"
  exit 1
fi

case $1 in
  main-build )
    # Run the tests first so the binary wont build if tests fail
    echo "Running NPC Generator tests"
    go test ./...

    echo "Building NPC Generator Binary"
    go build -o npcgen ./cmd
    ;;
  dev-build )
    # Run the tests first so the binary wont build if tests fail
    echo "Running NPC Generator Devtools tests"
    go test ./...

    echo "Building NPC Generator Binary"
    go build -o devtools ./cmd/devtools
    ;;
  release-build )
    # Run the tests first so the binary wont build if tests fail
    echo "Running NPC Generator tests"
    go test ./...

    echo "Building NPC Generator Binary"
    # Linux 64-bit
    GOOS=linux GOARCH=amd64 go build -o npcgen-linux

    # Windows 64-bit
    GOOS=windows GOARCH=amd64 go build -o npcgen.exe
    ;;
  
esac

echo "Build finished"