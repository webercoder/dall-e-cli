# Variables
BINARY_NAME=bin/dall-e-cli

# Default target
all: build

# Build the application
build:
	go build -o $(BINARY_NAME) .

# Clean up build artifacts
clean:
	rm -f $(BINARY_NAME)

# Run the application
run: build
	./$(BINARY_NAME) $(ARGS)

.PHONY: all build clean run
