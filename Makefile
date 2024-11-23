SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

BIN_NAME := musiclib

# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

.PHONY: run
#? run: Run the application
run:
	@go run cmd/$(BIN_NAME)/$(BIN_NAME).go

#? dist: Create the "dist" directory
dist:
	mkdir -p dist

.PHONY: binary
#? binary: Build the binary
binary: dist
	CGO_ENABLED=0 GOGC=off GOOS=${GOOS} GOARCH=${GOARCH} go build ${FLAGS[*]} -ldflags -s \
    -o "./dist/${GOOS}/${GOARCH}/$(BIN_NAME)" ./cmd/$(BIN_NAME)

.PHONY: test
#? test: Run the unit and integration tests
test: test-unit test-integration

.PHONY: test-unit
#? test-unit: Run the unit tests
test-unit:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go test -cover "-coverprofile=cover.out" -v $(TESTFLAGS) ./pkg/... ./cmd/... ./internal/...

.PHONY: test-integration
#? test-integration: Run the integration tests
test-integration: binary
	GOOS=$(GOOS) GOARCH=$(GOARCH) go test ./integration -test.timeout=20m -failfast -v $(TESTFLAGS)

.PHONY: lint
#? lint: Run golangci-lint
lint:
	golangci-lint run

.PHONY: fmt
#? fmt: Format the Code
fmt:
	gofmt -s -l -w $(SRCS)

.PHONY: help
#? help: Get more info on make commands
help: Makefile
	@echo " Choose a command run:"
	@sed -n 's/^#?//p' $< | column -t -s ':' |  sort | sed -e 's/^/ /'
