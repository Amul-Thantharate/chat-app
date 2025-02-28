# Go parameters
APP_NAME = chat-server
SRC = main.go
PORT = 8080

# Detect OS
OS := $(shell uname -s)

# Build the application
build:
ifeq ($(OS), Windows_NT)
	@echo "Building $(APP_NAME) for Windows... 🏗️"
	@set GOOS=windows && set GOARCH=amd64 && go build -o $(APP_NAME).exe $(SRC)
else
	@echo "Building $(APP_NAME) for Linux/Mac... 🏗️"
	@GOOS=linux GOARCH=amd64 go build -o $(APP_NAME) $(SRC)
endif

# Run the application
run: build
ifeq ($(OS), Windows_NT)
	@echo "Running $(APP_NAME) on port $(PORT)... 🚀"
	@$(APP_NAME).exe -addr :$(PORT)
else
	@echo "Running $(APP_NAME) on port $(PORT)... 🚀"
	@./$(APP_NAME) -addr :$(PORT)
endif

# Clean up generated files
clean:
ifeq ($(OS), Windows_NT)
	@echo "Cleaning up... 🧼"
	@if exist $(APP_NAME).exe del $(APP_NAME).exe
else
	@echo "Cleaning up... 🧼"
	@rm -f $(APP_NAME)
endif

# Format Go code
fmt:
	@echo "Formatting code... ✨"
	@go fmt ./...

# Lint the code (requires golangci-lint)
lint:
	@echo "Running linter... 🚦"
	@golangci-lint run

# Run tests
test:
	@echo "Running tests... 🧪"
	@go test ./...

# Help menu
help:
	@echo "Makefile commands: 📜"
	@echo "  build   - Compile the application 🔨"
	@echo "  run     - Build and run the server 🚀"
	@echo "  clean   - Remove built files 🧹"
	@echo "  fmt     - Format Go source files 🎨"
	@echo "  lint    - Run linter 🧐"
	@echo "  test    - Run tests ✅"

.PHONY: build run clean fmt lint test help
