all: build test

build:
	@echo "Building..."
	
	@go build -o main cmd/musiclib/main.go

run:
	@go run cmd/musiclib/main.go

test:
	@echo "Testing..."
	@go test ./... -v
