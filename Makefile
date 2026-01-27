GOCMD = go
GOBUILD = $(GOCMD) build

BINARY_NAME = felarof
BUILD_DIR = bin
MAIN_PACKAGE = ./core

all: build

deps:
	@echo "Checking Go dependencies..."
	@cd core && if [ ! -f go.mod ]; then \
		echo "Initializing Go module..."; \
		$(GOCMD) mod init $(BINARY_NAME); \
	fi
	@cd core && $(GOCMD) mod tidy

build: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME): deps
	@echo "Building Go backend..."
	mkdir -p $(BUILD_DIR)
	cd core && $(GOBUILD) -ldflags="-s -w" -o ../$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

.PHONY: all build deps clean
