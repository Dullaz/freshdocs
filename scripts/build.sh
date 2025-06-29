#!/bin/bash

set -e

# Get version from git tag or default to dev
VERSION=${1:-$(git describe --tags --always --dirty)}
echo "Building FreshDocs version: $VERSION"

# Create dist directory
mkdir -p dist

# Build for multiple platforms
PLATFORMS=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
)

for platform in "${PLATFORMS[@]}"; do
    IFS='/' read -r GOOS GOARCH <<< "$platform"
    
    echo "Building for $GOOS/$GOARCH..."
    
    # Set environment variables for cross-compilation
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    export CGO_ENABLED=0
    
    # Build the binary
    go build -ldflags="-s -w -X main.version=$VERSION" -o "dist/freshdocs-$GOOS-$GOARCH" .
    
    # Create archive
    if [ "$GOOS" = "windows" ]; then
        zip "dist/freshdocs-$VERSION-$GOOS-$GOARCH.zip" "dist/freshdocs-$GOOS-$GOARCH"
    else
        tar -czf "dist/freshdocs-$VERSION-$GOOS-$GOARCH.tar.gz" -C dist "freshdocs-$GOOS-$GOARCH"
    fi
done

echo "Build complete! Binaries are in the dist/ directory." 