.PHONY: build build-all build-darwin build-windows build-linux clean test

# Default target
build:
	go build -o unipm .

# Build for all platforms
build-all: build-darwin build-windows build-linux

# Build for macOS
build-darwin:
	@echo "Building for macOS (amd64)..."
	GOOS=darwin GOARCH=amd64 go build -o dist/unipm-darwin-amd64 .
	@echo "Building for macOS (arm64)..."
	GOOS=darwin GOARCH=arm64 go build -o dist/unipm-darwin-arm64 .

# Build for Windows
build-windows:
	@echo "Building for Windows (amd64)..."
	GOOS=windows GOARCH=amd64 go build -o dist/unipm-windows-amd64.exe .

# Build for Linux
build-linux:
	@echo "Building for Linux (amd64)..."
	GOOS=linux GOARCH=amd64 go build -o dist/unipm-linux-amd64 .
	@echo "Building for Linux (arm64)..."
	GOOS=linux GOARCH=arm64 go build -o dist/unipm-linux-arm64 .

# Clean build artifacts
clean:
	rm -rf dist/
	rm -f unipm unipm.exe

# Run tests
test:
	go test ./...

# Install locally
install:
	go install .
