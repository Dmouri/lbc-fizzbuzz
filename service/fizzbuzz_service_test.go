package service_test

import (
	"lbc/fizzbuzz/domain"
	"lbc/fizzbuzz/internal"
	"lbc/fizzbuzz/repository"
	"lbc/fizzbuzz/service"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// TestGenerateFizzBuzz /
func TestGenerateFizzBuzz(t *testing.T) {
	tests := []struct {
		name      string
		input     domain.FizzBuzzInput
		expected  string
		expectErr bool
	}{
		{
			name:      "Basic FizzBuzz",
			input:     domain.FizzBuzzInput{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
			expected:  "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz",
			expectErr: false,
		},
		{
			name:      "No multiples",
			input:     domain.FizzBuzzInput{Int1: 7, Int2: 8, Limit: 5, Str1: "fizz", Str2: "buzz"},
			expected:  "1,2,3,4,5",
			expectErr: false,
		},
		{
			name:      "Only multiples of Int1",
			input:     domain.FizzBuzzInput{Int1: 2, Int2: 10, Limit: 5, Str1: "fizz", Str2: "buzz"},
			expected:  "1,fizz,3,fizz,5",
			expectErr: false,
		},
		{
			name:      "Invalid Int1 and Int2",
			input:     domain.FizzBuzzInput{Int1: 0, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"},
			expectErr: true,
		},
		{
			name:      "Different strings",
			input:     domain.FizzBuzzInput{Int1: 2, Int2: 3, Limit: 6, Str1: "foo", Str2: "bar"},
			expected:  "1,foo,bar,foo,5,foobar",
			expectErr: false,
		},
		{
			name:      "Negative integers",
			input:     domain.FizzBuzzInput{Int1: -3, Int2: -5, Limit: 15, Str1: "fizz", Str2: "buzz"},
			expected:  "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz",
			expectErr: false,
		},
		{
			name:      "Empty strings",
			input:     domain.FizzBuzzInput{Int1: 3, Int2: 5, Limit: 15, Str1: "", Str2: "buzz"},
			expectErr: true,
		},
		{
			name:      "Negative limit",
			input:     domain.FizzBuzzInput{Int1: 3, Int2: 5, Limit: -5, Str1: "fizz", Str2: "buzz"},
			expectErr: true,
		},
		{
			name:      "Only multiples of Int1",
			input:     domain.FizzBuzzInput{Int1: 2, Int2: 10, Limit: 5, Str1: "fizz", Str2: "buzz"},
			expected:  "1,fizz,3,fizz,5",
			expectErr: false,
		},
		{
			name:      "Negative limit",
			input:     domain.FizzBuzzInput{Int1: 3, Int2: 5, Limit: -5, Str1: "fizz", Str2: "buzz"},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewFizzBuzzRepository(internal.Clients.PostgreSQL(), zap.NewExample())
			svc := service.NewFizzBuzzService(repo)
			result, err := svc.GenerateFizzBuzz(tt.input)

			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// TestGenerateFizzBuzz_LargeLimit /
func TestGenerateFizzBuzz_LargeLimit(t *testing.T) {
	repo := repository.NewFizzBuzzRepository(internal.Clients.PostgreSQL(), zap.NewExample())
	svc := service.NewFizzBuzzService(repo)
	tests := []struct {
		name          string
		input         domain.FizzBuzzInput
		expectedStart string
		expectedEnd   string
	}{
		{
			name: "Limit 1000",
			input: domain.FizzBuzzInput{
				Int1:  3,
				Int2:  5,
				Limit: 1000,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedStart: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz",
			expectedEnd:   "991,992,fizz,994,buzz,fizz,997,998,fizz,buzz",
		},
		{
			name: "Limit 10000",
			input: domain.FizzBuzzInput{
				Int1:  3,
				Int2:  5,
				Limit: 10000,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedStart: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz",
			expectedEnd:   "9991,9992,fizz,9994,buzz,fizz,9997,9998,fizz,buzz",
		},
		{
			name: "Limit 100000",
			input: domain.FizzBuzzInput{
				Int1:  3,
				Int2:  5,
				Limit: 100000,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedStart: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz",
			expectedEnd:   "99991,99992,fizz,99994,buzz,fizz,99997,99998,fizz,buzz",
		},
		{
			name: "Large int1 and int2 values",
			input: domain.FizzBuzzInput{
				Int1:  99999,
				Int2:  88888,
				Limit: 1000000,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedStart: "1,2,3,4,5,6,7,8,9,10",
			expectedEnd:   "999991,999992,999993,999994,999995,999996,999997,999998,999999,1000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := svc.GenerateFizzBuzz(tt.input)
			require.NoError(t, err)
			resultSlice := strings.Split(result, ",")
			assert.Equal(t, tt.expectedStart, strings.Join(resultSlice[:10], ","))
			assert.Equal(t, tt.expectedEnd, strings.Join(resultSlice[len(resultSlice)-10:], ","))
		})
	}
}
