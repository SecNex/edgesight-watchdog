#!/bin/bash

NAME="edgesight-watchdog"

# Build the project for windows (x64), linux (arm64, x64) and macos (arm64, x64)
echo "Building the project for windows (x64), linux (arm64, x64) and macos (arm64, x64)..."

# Create build directories if not exists
mkdir -p build/{windows/amd64,linux/{amd64,arm64},darwin/{amd64,arm64}}

# Windows
GOOS=windows GOARCH=amd64 go build -o build/windows/amd64/$NAME.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/$NAME
GOOS=linux GOARCH=arm64 go build -o build/linux/arm64/$NAME

# MacOS
GOOS=darwin GOARCH=amd64 go build -o build/darwin/amd64/$NAME
GOOS=darwin GOARCH=arm64 go build -o build/darwin/arm64/$NAME

# Done
echo "Build completed!"