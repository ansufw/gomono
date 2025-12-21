.PHONY: help build run stop clean dev generate

# Default target
help:
	@echo "Available commands:"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make dev      - Run in development mode with hot reload"
	@echo "  make stop     - Stop the running application"
	@echo "  make generate - Generate templates"
	@echo "  make clean    - Clean build files"

# Build the application
build:
	@echo "Building application..."
	@go build -o bin/web cmd/*.go

# Run the application
run:
	@echo "Starting application on http://localhost:3000"
	@go run cmd/web/main.go

# Run in development mode with hot reload
dev:
	@echo "Starting in development mode with hot reload..."
	@templ generate --watch &
	@go run cmd/web/main.go

# Stop the application
stop:
	@echo "Stopping application..."
	@pkill -f "go run cmd/web/main.go" || true
	@pkill -f "templ generate --watch" || true

# Generate templates
generate:
	@echo "Generating templates..."
	@templ generate

# Clean build files
clean:
	@echo "Cleaning build files..."
	@rm -rf bin/
	@go clean