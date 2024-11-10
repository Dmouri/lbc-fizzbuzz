# Variables
APP_NAME = fizzbuzz-api
GO = go
GOTEST = $(GO) test
GOMOD = $(GO) mod
TEST_COVERAGE_OUT = coverage.out

deps:
	$(GOMOD) tidy

build: deps
	$(GO) build -o $(APP_NAME) main.go

run: deps
	$(GO) run main.go

lint:
	golangci-lint run

test:
	$(GOTEST) ./... -v -coverprofile=$(TEST_COVERAGE_OUT)
	$(GO) tool cover -func=$(TEST_COVERAGE_OUT)

clean:
	rm -f $(APP_NAME) $(TEST_COVERAGE_OUT)

.PHONY: deps build run lint test clean swagger ci
