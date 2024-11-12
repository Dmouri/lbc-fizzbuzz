# lbc/fizzbuzz-api

A web server in Go that exposes a REST API for generating a customizable FizzBuzz sequence. The API allows setting specific values for sequence replacements.

## Prerequisites
- Go 1.22 or later
- golangci-lint (required for linting)
- Docker (for running PostgreSQL database)

## Installation

Clone the repository and initialize dependencies:

```sh
git clone https://github.com/Dmouri/lbc-fizzbuzz
cd lbc-fizzbuzz
go mod tidy
```

## Database Setup
This project uses a PostgreSQL database for persisting FizzBuzz requests. Use Docker Compose to set up and initialize the database.

- **Start the database**:
```sh
  make db-up
```

- **Initialize the database schema**:
  This runs the `init.sql` script to create the necessary tables.
```sh
  make db-init
```

- **Stop the database**:
```sh
  make db-down
```


## Usage

You can use the provided Makefile to manage building, running, testing, and linting the application:

- **Build the application**:
```sh
  make build
```

- **Run the application** (includes starting the database if not already running):
```sh
  make run
```

- **Run tests with coverage**:
```sh
  make test
```

- **Lint the code**:
  Ensure `golangci-lint` is installed, then run:
```sh
  make lint
```

- **Clean up build files**:
```sh
  make clean
```

## API Endpoints

### Custom FizzBuzz Sequence

This endpoint generates a customizable FizzBuzz sequence.

- **Endpoint**: `GET /api/v1/fizzbuzz`
- **Parameters**:
  - `int1`: integer to replace multiples of this value with `str1`
  - `int2`: integer to replace multiples of this value with `str2`
  - `limit`: upper limit of the sequence (e.g., 100)
  - `str1`: word that replaces multiples of `int1`
  - `str2`: word that replaces multiples of `int2`

Example:
```sh
curl "http://localhost:8080/api/v1/fizzbuzz?int1=3&int2=5&limit=16&str1=fizz&str2=buzz"
```

**Expected Output**:
```
{
  "result": "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16"
}
```

### FizzBuzz Statistics

This endpoint retrieves the FizzBuzz query that has been requested the most, displaying the parameters with the highest number of hits.

- **Endpoint**: `GET /api/v1/fizzbuzz/stats`
- **Response**:
  - `int1`: the integer used to replace multiples with `str1`
  - `int2`: the integer used to replace multiples with `str2`
  - `limit`: the upper limit of the sequence
  - `str1`: the word replacing multiples of `int1`
  - `str2`: the word replacing multiples of `int2`
  - `hits`: the number of times this specific query has been requested

**Example Response**:
```json
{
  "int1": 3,
  "int2": 5,
  "limit": 100,
  "str1": "fizz",
  "str2": "buzz",
  "hits": 42
}
```

This endpoint allows you to track the most popular FizzBuzz query configurations and observe usage patterns based on request frequency.
