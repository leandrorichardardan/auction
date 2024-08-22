.PHONY: all tidy test fmt

# Run all tasks
all: tidy fmt test

# Tidy the Go module dependencies
tidy:
	go mod tidy

# Format the Go code
fmt:
	go fmt ./...

# Run unit tests
test:
	go test ./... -v