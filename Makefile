all: build test

build:
	@echo "Building..."
	
	@go build -o main cmd/musiclib-api/main.go

run:
	@go run cmd/musiclib-api/main.go

test:
	@echo "Testing..."
	@go test ./... -v
