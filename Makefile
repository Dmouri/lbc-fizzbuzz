# Variables
APP_NAME = fizzbuzz-api
DB_CONTAINER = db
GO = go
GOTEST = $(GO) test
GOMOD = $(GO) mod
TEST_COVERAGE_OUT = coverage.out

db-up:
	docker-compose up -d $(DB_CONTAINER)

db-init: db-up
	docker exec -i $(DB_CONTAINER) psql -U fizzbuzz -d fizzbuzz_db < init.sql

db-down:
	docker-compose down

deps:
	$(GOMOD) tidy

build: deps
	$(GO) build -o $(APP_NAME) main.go

run: deps db-up
	$(GO) run main.go

lint:
	golangci-lint run ./api/... ./service/... ./repository/... --skip-dirs=/opt/hostedtoolcache

test:
	$(GOTEST) ./... -v -coverprofile=$(TEST_COVERAGE_OUT)
	$(GO) tool cover -func=$(TEST_COVERAGE_OUT)

clean:
	rm -f $(APP_NAME) $(TEST_COVERAGE_OUT)

.PHONY: deps build run lint test clean db-up db-init db-down
