APP_NAME=sticky
CMD_DIR=./cmd/sticky

.PHONY: run build install clean

# Run the app without building
run:
	go run $(CMD_DIR)

# Build the binary locally
build:
	go build -o bin/$(APP_NAME) $(CMD_DIR)

# Install into GOPATH/bin
install:
	go install $(CMD_DIR)

# Clean build artifacts
clean:
	rm -rf bin
