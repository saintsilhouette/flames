COVERAGE_FILE ?= coverage.out
TARGET ?= flame # CHANGE THIS TO YOUR BINARY NAME/NAMES

## build: build application
.PHONY: build
build:
	@echo "go build for the target ${TARGET} is in progress"
	@mkdir -p ./bin
	@go build -o ./bin/${TARGET} ./cmd/${TARGET}

## test: run all tests
.PHONY: test
test:
	@go test -coverpkg='github.com/es-debug/backend-academy-2024-go-template/...' --race -count=1 -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'

## run: build the target if necessary, then run it
.PHONY: run
run: build
	@echo "Running ./bin/${TARGET}"
	@./bin/${TARGET}

## clear: remove generated binaries and the 'images' directory
.PHONY: clear
clear:
	@echo "Removing build files (./bin/${TARGET}) and the 'images' directory if it exists..."
	@rm -rf bin images coverage.out