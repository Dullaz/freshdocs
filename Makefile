.PHONY: build clean test release help

# Default target
help:
	@echo "Available targets:"
	@echo "  build    - Build the freshdocs binary"
	@echo "  clean    - Clean build artifacts"
	@echo "  test     - Run tests"
	@echo "  release  - Build release binaries"
	@echo "  install  - Install freshdocs locally"

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