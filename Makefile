all: build test

.PHONY: lint
#? lint: Run golangci-lint
lint:
	golangci-lint run

build:
	@echo "Building..."
	
	@go build -o main cmd/musiclib/musiclib.go

run:
	@go run cmd/musiclib/musiclib.go

test:
	@echo "Testing..."
	@go test ./... -v
