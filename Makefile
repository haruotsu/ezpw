.PHONY: build test clean install lint fmt

# Binary name and version
BINARY_NAME=ezpw
VERSION?=v0.0.1

# Build binary
build:
	go build -ldflags="-X main.Version=$(VERSION)" -o bin/$(BINARY_NAME) cmd/ezpw/main.go

# Cross-compile for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -ldflags="-X main.Version=$(VERSION)" -o bin/$(BINARY_NAME)-linux-amd64 cmd/ezpw/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.Version=$(VERSION)" -o bin/$(BINARY_NAME)-darwin-amd64 cmd/ezpw/main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags="-X main.Version=$(VERSION)" -o bin/$(BINARY_NAME)-darwin-arm64 cmd/ezpw/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-X main.Version=$(VERSION)" -o bin/$(BINARY_NAME)-windows-amd64.exe cmd/ezpw/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	gofmt -s -w .
	go mod tidy

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install binary
install:
	go install -ldflags="-X main.Version=$(VERSION)" cmd/ezpw/main.go

# Development build (format, lint, test, build)
dev: fmt lint test build

# Release build (clean, format, lint, test, build-all)
release: clean fmt lint test build-all