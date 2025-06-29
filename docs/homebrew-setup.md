# Homebrew Setup Guide

This document explains how to set up FreshDocs for distribution via Homebrew using a separate tap repository.

## Prerequisites

1. A GitHub repository for your project (this repository)
2. A separate GitHub repository for your Homebrew tap: `dullaz/homebrew-freshdocs`

## Setup Steps

### 1. Create the Homebrew Tap Repository

Create a new GitHub repository named `homebrew-freshdocs` with the following structure:

```
homebrew-freshdocs/
└── Formula/
    └── freshdocs.rb
```

### 2. Set up GitHub Secrets

In your main repository (`dullaz/freshdocs`), go to Settings > Secrets and variables > Actions and add:

- `HOMEBREW_TAP_TOKEN`: A GitHub Personal Access Token with repo permissions for the `dullaz/homebrew-freshdocs` repository

### 3. Release Process

1. **Tag a release**: `git tag v1.0.0 && git push origin v1.0.0`
2. **GitHub Actions** will automatically:
   - Build binaries for all platforms
   - Create a GitHub release with the binaries
   - Update the Homebrew formula in `dullaz/homebrew-freshdocs`

### 4. Install via Homebrew

Users can then install your tool with:

```bash
# Add your tap
brew tap dullaz/freshdocs

# Install freshdocs
brew install freshdocs
```

## Manual Release Process

If you prefer to release manually:

1. Run the build script: `./scripts/build.sh v1.0.0`
2. Create a GitHub release with the generated files
3. Manually update the `Formula/freshdocs.rb` file in `dullaz/homebrew-freshdocs`

## Example Formula

The workflow will create a formula like this in `dullaz/homebrew-freshdocs/Formula/freshdocs.rb`:

```ruby
class Freshdocs < Formula
  desc "Keep your documentation as fresh as your code"
  homepage "https://github.com/dullaz/freshdocs"
  version "v1.0.0"
  
  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdocs/releases/download/v1.0.0/freshdocs-v1.0.0-darwin-arm64.tar.gz"
      sha256 "your-sha256-here"
    else
      url "https://github.com/dullaz/freshdocs/releases/download/v1.0.0/freshdocs-v1.0.0-darwin-amd64.tar.gz"
      sha256 "your-sha256-here"
    end
  end
  
  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdocs/releases/download/v1.0.0/freshdocs-v1.0.0-linux-arm64.tar.gz"
      sha256 "your-sha256-here"
    else
      url "https://github.com/dullaz/freshdocs/releases/download/v1.0.0/freshdocs-v1.0.0-linux-amd64.tar.gz"
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

- **SHA256 mismatch**: The workflow automatically calculates the correct hashes from the released binaries
- **URL errors**: Verify the release URLs are correct and accessible
- **Build failures**: Check that all dependencies are properly specified in `go.mod`
- **Permission errors**: Ensure the `HOMEBREW_TAP_TOKEN` has write access to `dullaz/homebrew-freshdocs` 