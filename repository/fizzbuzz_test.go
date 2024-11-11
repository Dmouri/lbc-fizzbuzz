package repository

import (
	"context"
	"lbc/fizzbuzz/domain"
	"lbc/fizzbuzz/internal"
	"lbc/fizzbuzz/testdata/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestFizzBuzzRepositorySave(t *testing.T) {
	db := internal.Clients.PostgreSQL()
	logger := zap.NewExample()
	repo := NewFizzBuzzRepository(db, logger)

	err := utils.ResetDatabase(db)
	assert.Nil(t, err)

	tests := []struct {
		name         string
		input        domain.FizzBuzzInput
		expectedHits int
		runSaveTwice bool
	}{
		{
			name: "Create New Entry",
			input: domain.FizzBuzzInput{
				Int1:  3,
				Int2:  5,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedHits: 1,
		},
		{
			name: "Increment Hits for Existing Entry",
			input: domain.FizzBuzzInput{
				Int1:  3,
				Int2:  7,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedHits: 2,
			runSaveTwice: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.runSaveTwice {
				err := repo.Save(context.Background(), tt.input)
				assert.Nil(t, err)
			}

			err := repo.Save(context.Background(), tt.input)
			assert.Nil(t, err)

			var result domain.FizzbuzzRequest
			errSQL := db.NewSelect().
				Model(&result).
				Where("int1 = ? AND int2 = ? AND max_limit = ? AND str1 = ? AND str2 = ?", tt.input.Int1, tt.input.Int2, tt.input.Limit, tt.input.Str1, tt.input.Str2).
				Scan(context.Background())
			assert.Nil(t, errSQL)
			assert.Equal(t, tt.expectedHits, result.Hits)
		})
	}
}

func TestFizzBuzzRepositoryGetMostHits(t *testing.T) {
	db := internal.Clients.PostgreSQL()
	logger := zap.NewExample()
	repo := NewFizzBuzzRepository(db, logger)

	err := utils.ResetDatabase(db)
	assert.Nil(t, err)

	tests := []struct {
		name           string
		setup          func()
		expectedResult domain.FizzbuzzRequest
	}{
		{
			name: "Get Entry with Most Hits",
			setup: func() {
				input1 := domain.FizzBuzzInput{
					Int1:  3,
					Int2:  5,
					Limit: 100,
					Str1:  "fizz",
					Str2:  "buzz",
				}
				input2 := domain.FizzBuzzInput{
					Int1:  2,
					Int2:  7,
					Limit: 50,
					Str1:  "foo",
					Str2:  "bar",
				}

				errSQL := repo.Save(context.Background(), input1)
				assert.Nil(t, errSQL)
				errSQL = repo.Save(context.Background(), input1)
				assert.Nil(t, errSQL)
				errSQL = repo.Save(context.Background(), input2)
				assert.Nil(t, errSQL)
			},
			expectedResult: domain.FizzbuzzRequest{
				FizzBuzzInput: domain.FizzBuzzInput{
					Int1:  3,
					Int2:  5,
					Limit: 100,
					Str1:  "fizz",
					Str2:  "buzz",
				},
				Hits: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			result, err := repo.GetMostHits(context.Background())
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedResult.Int1, result.Int1)
			assert.Equal(t, tt.expectedResult.Int2, result.Int2)
			assert.Equal(t, tt.expectedResult.Limit, result.Limit)
			assert.Equal(t, tt.expectedResult.Str1, result.Str1)
			assert.Equal(t, tt.expectedResult.Str2, result.Str2)
			assert.Equal(t, tt.expectedResult.Hits, result.Hits)
		})
	}
}
