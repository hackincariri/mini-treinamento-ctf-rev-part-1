#!/bin/bash

# Output file name
OUTPUT_NAME="hello_world"

# Output directory
BUILD_DIR="build"
mkdir -p $BUILD_DIR

# Compile for Windows
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/${OUTPUT_NAME}.exe main.go

# Compile for Linux
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/${OUTPUT_NAME}_linux main.go

# Compile for macOS
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/${OUTPUT_NAME}_mac main.go

# Compile for (arm64)
GOOS=darwin GOARCH=arm64 go build -o $BUILD_DIR/${OUTPUT_NAME}_mac_arm main.go

# Success message
echo "Build finished! Binary files are in ./$BUILD_DIR/"