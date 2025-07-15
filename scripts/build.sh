#!/bin/bash

if [ -z $1 ]; then
  echo "Must provide build mode as argument"
  echo "Options are:"
  echo "  - main-build          Builds and runs the main binary"
  echo "  - main-test           Builds the main binary and runs its tests"
  echo "  - dev-build           Builds and runs the devtools binary"
  echo "  - dev-test            Builds the devtools binary and runs its tests"
  exit 1
fi

case $1 in
  main-build )
    echo "Building NPC Generator Binary"
    go build -o npcgen ./cmd
    ;;
  main-test )
    # Run the tests first so the binary wont build if tests fail
    echo "Running NPC Generator tests"
    go test ./...

    echo "Building NPC Generator Binary"
    go build -o npcgen ./cmd
    ;;
  dev-build )
    echo "Building NPC Generator Binary"
    go build -o devtools ./cmd/devtools
    ;;
  dev-test )
    # Run the tests first so the binary wont build if tests fail
    echo "Running NPC Generator Devtools tests"
    go test ./...

    echo "Building NPC Generator Devtools Binary"
    go build -o devtools ./cmd/devtools
    ;;
  
esac

echo "Build finished"