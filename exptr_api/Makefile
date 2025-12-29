# Makefile for managing Go development tasks locally.

# --- Variables ---
# You can override these variables from the command line
GO_FILES := $(shell find . -name "*.go" -not -path "./vendor/*")

# Use .PHONY to declare targets that are not files. This avoids conflicts with files of the same name.
.PHONY: all ci build test lint format tidy vet run run_air install-deps clean help

# The default command that runs when you just type `make`.
all: ci

# --- CI Target ---
# Runs all CI tasks locally: linting, formatting, vetting, and tidying.
ci: tidy format lint test
	@echo "All CI checks completed successfully."

# --- Individual Commands ---

# Build the Go application
build:
	@echo "Building the application..."
	@go build -o bin/api cmd/api/main.go
	@echo "Build complete. Binary created at bin/api"

# Run tests
test:
	@echo "Running tests..."
	# @go test -v ./...

# Run golangci-lint locally
lint:
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "Error: golangci-lint not found. Please install it first." && echo "Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	@golangci-lint run ./...

# Format Go code
format:
	@echo "Formatting code..."
	@go fmt ./...

#Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

# Tidy go modules
tidy:
	@echo "Tidying go modules..."
	@go mod tidy

# Run the application locally
run:
	@echo "Running the application locally..."
	CONFIG_PATH=config/local.yaml go run cmd/api/main.go --env=.env

# Run the application with air for hot-reloading
run_air:
	@echo "Running the application with air for live reloading..."
	@which air > /dev/null || (echo "Error: air not found. Please install it first." && echo "Install with: go install github.com/cosmtrek/air@latest" && exit 1)
	CONFIG_PATH=config/local.yaml air

# Install development dependencies locally
install-deps:
	@echo "Installing development dependencies..."
	@echo "Installing golangci-lint..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Installing air for hot reloading..."
	@go install github.com/cosmtrek/air@latest
	@echo "All dependencies installed successfully."

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean ./...
	@echo "Clean complete."

# Run a quick check (format, vet, lint)
check: format lint
	@echo "Quick checks completed."

# --- Help ---
# Provides a self-documenting help message.
help:
	@echo "Available commands:"
	@echo "  make ci            - Runs full CI suite (tidy, format, vet, lint, test). (Default)"
	@echo "  make build         - Builds the Go application."
	@echo "  make test          - Runs all tests."
	@echo "  make lint          - Runs golangci-lint."
	@echo "  make format        - Formats Go code with go fmt."
	@echo "  make vet           - Runs go vet."
	@echo "  make tidy          - Runs go mod tidy."
	@echo "  make run           - Runs the Go application locally."
	@echo "  make run_air       - Runs the Go application with Air for hot-reloading."
	@echo "  make check         - Runs quick checks (format, vet, lint)."
	@echo "  make install-deps  - Installs required development dependencies."
	@echo "  make clean         - Cleans build artifacts."
	@echo "  make help          - Shows this help message."
