.PHONY: build clean run watch docker-build docker-run docker-stop docker-logs compose-up compose-down

# Variables
IMAGE_NAME     ?= memap
TAG            ?= latest
CONTAINER_NAME ?= memap
PORT           ?= 2118

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

# Run the application locally
run:
	@go run ./cmd

# Live Reload
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

docker-build:
	@echo "Building docker image $(IMAGE_NAME):$(TAG)..."
	@docker build -t $(IMAGE_NAME):$(TAG) .

docker-run:
	@echo "Running container $(CONTAINER_NAME)..."
	@docker run -d --name $(CONTAINER_NAME) -p $(PORT):$(PORT) --restart unless-stopped $(IMAGE_NAME):$(TAG)

docker-stop:
	@echo "Stopping and removing container $(CONTAINER_NAME)..."
	@docker stop $(CONTAINER_NAME) 2>/dev/null || true
	@docker rm $(CONTAINER_NAME) 2>/dev/null || true

docker-logs:
	@docker logs -f $(CONTAINER_NAME)

# Docker Compose
compose-up:
	@docker compose up -d

compose-down:
	@docker compose down
