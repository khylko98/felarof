BINARY_NAME = felarof
BUILD_DIR = bin

all: build

deps:
	@echo "Tidying dependencies..."
	go mod tidy

build: deps
	@echo "Building..."
	mkdir -p $(BUILD_DIR)
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	rm -rf $(BUILD_DIR)

.PHONY: all build deps clean
