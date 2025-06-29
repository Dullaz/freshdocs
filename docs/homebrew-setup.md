# Homebrew Setup Guide

This document explains how to set up FreshDocs for distribution via Homebrew.

## Prerequisites

1. A GitHub repository for your project
2. A separate GitHub repository for your Homebrew tap (e.g., `yourusername/homebrew-tap`)

## Setup Steps

### 1. Create the Homebrew Tap Repository

Create a new GitHub repository named `homebrew-tap` (or any name you prefer).

### 2. Set up GitHub Secrets

In your main repository, go to Settings > Secrets and variables > Actions and add:

- `HOMEBREW_TAP_TOKEN`: A GitHub Personal Access Token with repo permissions

### 3. Create the Formula Directory

In your `homebrew-tap` repository, create a `Formula/` directory and add a `freshdocs.rb` file:

```ruby
class Freshdocs < Formula
  desc "Keep your documentation as fresh as your code"
  homepage "https://github.com/yourusername/freshdocs"
  version "v1.0.0"
  
  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/yourusername/freshdocs/releases/download/v1.0.0/freshdocs-v1.0.0-darwin-arm64.tar.gz"
      sha256 "your-sha256-here"
    else
      url "https://github.com/yourusername/freshdocs/releases/download/v1.0.0/freshdocs-v1.0.0-darwin-amd64.tar.gz"
      sha256 "your-sha256-here"
    end
  end
  
  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/yourusername/freshdocs/releases/download/v1.0.0/freshdocs-v1.0.0-linux-arm64.tar.gz"
      sha256 "your-sha256-here"
    else
      url "https://github.com/yourusername/freshdocs/releases/download/v1.0.0/freshdocs-v1.0.0-linux-amd64.tar.gz"
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

### 4. Release Process

1. **Tag a release**: `git tag v1.0.0 && git push origin v1.0.0`
2. **GitHub Actions** will automatically:
   - Build binaries for all platforms
   - Create a GitHub release
   - Update the Homebrew formula

### 5. Install via Homebrew

Users can then install your tool with:

```bash
# Add your tap
brew tap yourusername/tap

# Install freshdocs
brew install freshdocs
```

## Manual Release Process

If you prefer to release manually:

1. Run the build script: `./scripts/build.sh v1.0.0`
2. Create a GitHub release with the generated files
3. Update the Homebrew formula manually with the correct URLs and SHA256 hashes

## Troubleshooting

- **SHA256 mismatch**: Make sure the SHA256 in the formula matches the actual file
- **URL errors**: Verify the release URLs are correct and accessible
- **Build failures**: Check that all dependencies are properly specified in `go.mod` 