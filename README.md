# lbc/fizzbuzz-api

A web server in Go that exposes a REST API for generating a customizable FizzBuzz sequence. The API allows setting specific values for sequence replacements.

## Prerequisites
- Go 1.22 or later

## Installation

Clone the repository and initialize dependencies:

```sh
git clone https://github.com/Dmouri/lbc-fizzbuzz
cd lbc-fizzbuzz
go mod tidy
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
