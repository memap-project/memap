.PHONY: build clean run watch

# Test the application without caching
test:
	@echo "Testing..."
	@go test -count=1 ./...

# Test with caching
test-cache:
	@echo "Testing with caching..."
	@go test ./...

# Build the application
build:
	@echo "Building..."
	@go build -o main ./cmd

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Run the application
run:
	@go run ./cmd

# Live Reload
# Locate air binary path
AIR_BIN := $(shell which air 2>/dev/null || echo "$(shell go env GOPATH)/bin/air")

watch:
	@if [ -x "$(AIR_BIN)" ]; then \
		echo "Starting air for live reload..."; \
		$(AIR_BIN); \
	else \
		echo "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] "; \
		read choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			echo "Installing air..."; \
			go install github.com/air-verse/air@latest; \
			echo "Starting air..."; \
			$(AIR_BIN); \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi
