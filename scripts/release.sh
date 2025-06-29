#!/bin/bash

set -e

if [ $# -eq 0 ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.0.0"
    exit 1
fi

VERSION=$1

echo "Creating release for version: $VERSION"

# Check if we're on main branch
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ]; then
    echo "Error: You must be on the main branch to create a release"
    exit 1
fi

# Check if working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    echo "Error: Working directory is not clean. Please commit or stash your changes."
    exit 1
fi

# Check if tag already exists
if git tag -l | grep -q "^$VERSION$"; then
    echo "Error: Tag $VERSION already exists"
    exit 1
fi

# Build binaries and create formula
echo "Building binaries and creating Homebrew formula..."
./scripts/build.sh "$VERSION"

# Add all files
git add .

# Commit changes
echo "Committing changes..."
git commit -m "Release $VERSION

- Build binaries for all platforms
- Update Homebrew formula"

# Create and push tag
echo "Creating and pushing tag..."
git tag "$VERSION"
git push origin main
git push origin "$VERSION"

echo "Release $VERSION created successfully!"
echo "GitHub Actions will now create the GitHub release with the binaries." 