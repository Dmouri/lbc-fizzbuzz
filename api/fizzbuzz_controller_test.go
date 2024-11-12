package api_test

import (
	"lbc/fizzbuzz/api"
	"lbc/fizzbuzz/internal"
	"lbc/fizzbuzz/repository"
	"lbc/fizzbuzz/service"
	"lbc/fizzbuzz/testdata/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// TestFizzBuzzEndpointValidation tests input validation for FizzBuzz API
func TestFizzBuzzEndpointValidation(t *testing.T) {
	router := gin.Default()
	fizzBuzzRepository := repository.NewFizzBuzzRepository(internal.Clients.PostgreSQL(), zap.NewExample())
	fizzBuzzService := service.NewFizzBuzzService(fizzBuzzRepository)
	api.SetupFizzBuzzController(zap.NewExample(), router, fizzBuzzService, fizzBuzzRepository)

	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid parameters",
			url:          "/api/v1/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz",
			expectedCode: http.StatusOK,
			expectedBody: `"result":"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz"`,
		},
		{
			name:         "Invalid int1 parameter (zero)",
			url:          "/api/v1/fizzbuzz?int1=0&int2=5&limit=15&str1=fizz&str2=buzz",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"message":"int1 must be different than 0","kind":"invalid_input"`,
		},
		{
			name:         "Invalid int2 parameter (zero)",
			url:          "/api/v1/fizzbuzz?int1=3&int2=0&limit=15&str1=fizz&str2=buzz",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"message":"int2 must be different than 0","kind":"invalid_input"`,
		},
		{
			name:         "Empty str1 parameter",
			url:          "/api/v1/fizzbuzz?int1=3&int2=5&limit=15&str1=&str2=buzz",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"message":"str1 must not be empty","kind":"invalid_input"`,
		},
		{
			name:         "Empty str2 parameter",
			url:          "/api/v1/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"message":"str2 must not be empty","kind":"invalid_input"`,
		},
		{
			name:         "Non-integer int1 parameter",
			url:          "/api/v1/fizzbuzz?int1=abc&int2=5&limit=15&str1=fizz&str2=buzz",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"message":"failed to parse int1","kind":"failed_to_parse_int1"`,
		},
		{
			name:         "Non-integer int2 parameter",
			url:          "/api/v1/fizzbuzz?int1=3&int2=xyz&limit=15&str1=fizz&str2=buzz",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"message":"failed to parse int2","kind":"failed_to_parse_int2"`,
		},
		{
			name:         "Non-integer limit parameter",
			url:          "/api/v1/fizzbuzz?int1=3&int2=5&limit=abc&str1=fizz&str2=buzz",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"message":"failed to parse limit","kind":"failed_to_parse_limit"`,
		},
		{
			name:         "Identical int1 and int2 parameters",
			url:          "/api/v1/fizzbuzz?int1=3&int2=3&limit=15&str1=fizz&str2=buzz",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"message":"int1 and int2 must be different","kind":"invalid_input"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}

// TestFizzBuzzEndpointResults tests the FizzBuzz API output for various valid inputs
func TestFizzBuzzEndpointResults(t *testing.T) {
	router := gin.Default()
	fizzBuzzRepository := repository.NewFizzBuzzRepository(internal.Clients.PostgreSQL(), zap.NewExample())
	fizzBuzzService := service.NewFizzBuzzService(fizzBuzzRepository)
	api.SetupFizzBuzzController(zap.NewExample(), router, fizzBuzzService, fizzBuzzRepository)

	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Standard FizzBuzz with limit 15",
			url:          "/api/v1/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz",
			expectedCode: http.StatusOK,
			expectedBody: `"result":"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz"`,
		},
		{
			name:         "Different strings for replacements",
			url:          "/api/v1/fizzbuzz?int1=3&int2=5&limit=15&str1=foo&str2=bar",
			expectedCode: http.StatusOK,
			expectedBody: `"result":"1,2,foo,4,bar,foo,7,8,foo,bar,11,foo,13,14,foobar"`,
		},
		{
			name:         "Higher limit",
			url:          "/api/v1/fizzbuzz?int1=3&int2=5&limit=20&str1=fizz&str2=buzz",
			expectedCode: http.StatusOK,
			expectedBody: `"result":"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz"`,
		},
		{
			name:         "Higher limit and different integers",
			url:          "/api/v1/fizzbuzz?int1=2&int2=7&limit=20&str1=fizz&str2=buzz",
			expectedCode: http.StatusOK,
			expectedBody: `"result":"1,fizz,3,fizz,5,fizz,buzz,fizz,9,fizz,11,fizz,13,fizzbuzz,15,fizz,17,fizz,19,fizz"`,
		},
		{
			name:         "Only multiples of int1",
			url:          "/api/v1/fizzbuzz?int1=2&int2=11&limit=10&str1=foo&str2=bar",
			expectedCode: http.StatusOK,
			expectedBody: `"result":"1,foo,3,foo,5,foo,7,foo,9,foo"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}

func TestGetFizzBuzzStatsEndpoint(t *testing.T) {
	router := gin.Default()
	db := internal.Clients.PostgreSQL()
	fizzBuzzRepository := repository.NewFizzBuzzRepository(db, zap.NewExample())
	fizzBuzzService := service.NewFizzBuzzService(fizzBuzzRepository)
	api.SetupFizzBuzzController(zap.NewExample(), router, fizzBuzzService, fizzBuzzRepository)

	err := utils.LoadFixtures(db)
	assert.Nil(t, err)

	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Successful Stats Fetch",
			url:          "/api/v1/fizzbuzz/stats",
			expectedCode: http.StatusOK,
			expectedBody: `"int1":3,"int2":5,"limit":100,"str1":"fizz","str2":"buzz","hits":42`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}
