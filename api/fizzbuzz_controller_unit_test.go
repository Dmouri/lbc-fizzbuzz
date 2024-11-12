package api_test

import (
	"lbc/fizzbuzz/api"
	"lbc/fizzbuzz/domain"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetQueryParams(t *testing.T) {
	tests := []struct {
		name      string
		query     map[string]string
		expected  domain.FizzBuzzInput
		expectErr bool
	}{
		{
			name: "Valid parameters",
			query: map[string]string{
				"int1": "3", "int2": "5", "limit": "15", "str1": "fizz", "str2": "buzz",
			},
			expected:  domain.FizzBuzzInput{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
			expectErr: false,
		},
		{
			name: "Missing limit - should use default",
			query: map[string]string{
				"int1": "3", "int2": "5", "str1": "fizz", "str2": "buzz",
			},
			expected:  domain.FizzBuzzInput{Int1: 3, Int2: 5, Limit: 100, Str1: "fizz", Str2: "buzz"},
			expectErr: false,
		},
		{
			name: "Negative int1",
			query: map[string]string{
				"int1": "-3", "int2": "5", "limit": "15", "str1": "fizz", "str2": "buzz",
			},
			expected:  domain.FizzBuzzInput{Int1: -3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
			expectErr: false,
		},
		{
			name: "Negative int2",
			query: map[string]string{
				"int1": "3", "int2": "-5", "limit": "15", "str1": "fizz", "str2": "buzz",
			},
			expected:  domain.FizzBuzzInput{Int1: 3, Int2: -5, Limit: 15, Str1: "fizz", Str2: "buzz"},
			expectErr: false,
		},
		{
			name: "Invalid int1 - non-integer value 'aaa'",
			query: map[string]string{
				"int1": "aaa", "int2": "5", "limit": "15", "str1": "fizz", "str2": "buzz",
			},
			expectErr: true,
		},
		{
			name: "Invalid int2 - non-integer value 'bbb'",
			query: map[string]string{
				"int1": "3", "int2": "bbb", "limit": "15", "str1": "fizz", "str2": "buzz",
			},
			expectErr: true,
		},
		{
			name: "Invalid limit - non-integer value 'ccc'",
			query: map[string]string{
				"int1": "3", "int2": "5", "limit": "ccc", "str1": "fizz", "str2": "buzz",
			},
			expectErr: true,
		},
		{
			name: "Negative limit - type is correct (will be rejected by validation later)",
			query: map[string]string{
				"int1": "3", "int2": "5", "limit": "-15", "str1": "fizz", "str2": "buzz",
			},
			expected:  domain.FizzBuzzInput{Int1: 3, Int2: 5, Limit: -15, Str1: "fizz", Str2: "buzz"},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			ctx := &gin.Context{}
			ctx.Request = &http.Request{
				URL: &url.URL{},
			}

			q := ctx.Request.URL.Query()
			for k, v := range tt.query {
				q.Add(k, v)
			}

			ctx.Request.URL.RawQuery = q.Encode()
			result, err := api.GetQueryParams(ctx)

			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}

		})
	}
}
