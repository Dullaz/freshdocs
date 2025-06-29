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

# Create Formula directory
mkdir -p Formula

# Calculate SHA256 hashes
DARWIN_AMD64_SHA=$(shasum -a 256 "dist/freshdocs-$VERSION-darwin-amd64.tar.gz" | cut -d' ' -f1)
DARWIN_ARM64_SHA=$(shasum -a 256 "dist/freshdocs-$VERSION-darwin-arm64.tar.gz" | cut -d' ' -f1)
LINUX_AMD64_SHA=$(shasum -a 256 "dist/freshdocs-$VERSION-linux-amd64.tar.gz" | cut -d' ' -f1)
LINUX_ARM64_SHA=$(shasum -a 256 "dist/freshdocs-$VERSION-linux-arm64.tar.gz" | cut -d' ' -f1)

# Create Homebrew formula
cat > Formula/freshdocs.rb << EOF
class Freshdocs < Formula
  desc "Keep your documentation as fresh as your code"
  homepage "https://github.com/dullaz/freshdoc"
  version "$VERSION"
  
  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdoc/releases/download/$VERSION/freshdocs-$VERSION-darwin-arm64.tar.gz"
      sha256 "$DARWIN_ARM64_SHA"
    else
      url "https://github.com/dullaz/freshdoc/releases/download/$VERSION/freshdocs-$VERSION-darwin-amd64.tar.gz"
      sha256 "$DARWIN_AMD64_SHA"
    end
  end
  
  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdoc/releases/download/$VERSION/freshdocs-$VERSION-linux-arm64.tar.gz"
      sha256 "$LINUX_ARM64_SHA"
    else
      url "https://github.com/dullaz/freshdoc/releases/download/$VERSION/freshdocs-$VERSION-linux-amd64.tar.gz"
      sha256 "$LINUX_AMD64_SHA"
    end
  end
  
  def install
    bin.install "freshdocs"
  end
  
  test do
    system "#{bin}/freshdocs", "--version"
  end
end
EOF

echo "Build complete! Binaries are in the dist/ directory."
echo "Homebrew formula created at Formula/freshdocs.rb" 