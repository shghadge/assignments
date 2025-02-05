#!/bin/bash

# List of target operating systems and architectures
targets=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
    "windows/arm64"
)

cd ./server
# Loop through each target and build the binary
for target in "${targets[@]}"; do
    OS=$(echo "$target" | cut -d/ -f1)
    ARCH=$(echo "$target" | cut -d/ -f2)

    echo "Building for $OS/$ARCH..."
    pwd
    GOOS="$OS" GOARCH="$ARCH" go build -o "server-binary-$OS-$ARCH"
done

cd ../
cd ./client/
for target in "${targets[@]}"; do
    OS=$(echo "$target" | cut -d/ -f1)
    ARCH=$(echo "$target" | cut -d/ -f2)

    echo "Building for $OS/$ARCH..."
    pwd
    GOOS="$OS" GOARCH="$ARCH" go build -o "client-binary-$OS-$ARCH"
done

