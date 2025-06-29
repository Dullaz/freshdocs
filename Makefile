.PHONY: build clean test release help

# Default target
help:
	@echo "Available targets:"
	@echo "  build    - Build the freshdocs binary"
	@echo "  clean    - Clean build artifacts"
	@echo "  test     - Run tests"
	@echo "  release  - Build release binaries"
	@echo "  install  - Install freshdocs locally"
	@echo "  new-release <version> - Create a new release (e.g., make new-release v1.0.0)"

# Build the binary
build:
	go build -o freshdocs .

# Clean build artifacts
clean:
	rm -rf freshdocs dist/

# Run tests
test:
	go test ./...

# Build release binaries
release:
	./scripts/build.sh

# Create a new release
new-release:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Usage: make new-release VERSION=v1.0.0"; \
		exit 1; \
	fi
	./scripts/release.sh $(VERSION)

# Install locally
install: build
	go install .

# Run the application
run: build
	./freshdocs

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Check for security issues
security:
	gosec ./... 