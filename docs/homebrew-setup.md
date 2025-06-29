# Homebrew Setup Guide

This document explains how to set up FreshDocs for distribution via Homebrew using the same repository.

## Prerequisites

1. A GitHub repository for your project (this repository)

## Setup Steps

### 1. The Formula Directory

The GitHub Actions workflow will automatically create a `Formula/` directory in your repository and add a `freshdocs.rb` file when you create a release.

### 2. Release Process

1. **Tag a release**: `git tag v1.0.0 && git push origin v1.0.0`
2. **GitHub Actions** will automatically:
   - Build binaries for all platforms
   - Create a GitHub release
   - Create/update the Homebrew formula in your repository

### 3. Install via Homebrew

Users can then install your tool with:

```bash
# Add your repository as a tap
brew tap yourusername/freshdoc

# Install freshdocs
brew install freshdocs
```

Or if you prefer a different tap name:

```bash
# Add your repository as a tap with a custom name
brew tap yourusername/freshdoc https://github.com/yourusername/freshdoc

# Install freshdocs
brew install freshdocs
```

## Manual Release Process

If you prefer to release manually:

1. Run the build script: `./scripts/build.sh v1.0.0`
2. Create a GitHub release with the generated files
3. Create the `Formula/freshdocs.rb` file manually with the correct URLs and SHA256 hashes

## Example Formula

The workflow will create a formula like this:

```ruby
class Freshdocs < Formula
  desc "Keep your documentation as fresh as your code"
  homepage "https://github.com/yourusername/freshdoc"
  version "v1.0.0"
  
  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/yourusername/freshdoc/releases/download/v1.0.0/freshdocs-v1.0.0-darwin-arm64.tar.gz"
      sha256 "your-sha256-here"
    else
      url "https://github.com/yourusername/freshdoc/releases/download/v1.0.0/freshdocs-v1.0.0-darwin-amd64.tar.gz"
      sha256 "your-sha256-here"
    end
  end
  
  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/yourusername/freshdoc/releases/download/v1.0.0/freshdocs-v1.0.0-linux-arm64.tar.gz"
      sha256 "your-sha256-here"
    else
      url "https://github.com/yourusername/freshdoc/releases/download/v1.0.0/freshdocs-v1.0.0-linux-amd64.tar.gz"
      sha256 "your-sha256-here"
    end
  end
  
  def install
    bin.install "freshdocs"
  end
  
  test do
    system "#{bin}/freshdocs", "--version"
  end
end
```

## Troubleshooting

- **SHA256 mismatch**: Make sure the SHA256 in the formula matches the actual file
- **URL errors**: Verify the release URLs are correct and accessible
- **Build failures**: Check that all dependencies are properly specified in `go.mod`
- **Permission errors**: Ensure the workflow has write permissions to the repository 