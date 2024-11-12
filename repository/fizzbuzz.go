package repository

import (
	"context"
	"lbc/fizzbuzz/domain"

	"github.com/mwm-io/gapi/errors"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type FizzBuzzRepository interface {
	Save(ctx context.Context, input domain.FizzBuzzInput) errors.Error
	GetMostHits(ctx context.Context) (domain.FizzbuzzRequest, errors.Error)
}

type fizzBuzzRepository struct {
	db     *bun.DB
	logger *zap.Logger
}

func NewFizzBuzzRepository(db *bun.DB, logger *zap.Logger) FizzBuzzRepository {
	return &fizzBuzzRepository{
		db:     db,
		logger: logger,
	}
}

func (f *fizzBuzzRepository) Save(ctx context.Context, input domain.FizzBuzzInput) errors.Error {
	_, err := f.db.NewInsert().
		Model(&domain.FizzbuzzRequest{
			FizzBuzzInput: input,
			Hits:          1,
		}).
		On("CONFLICT (int1, int2, max_limit, str1, str2) DO UPDATE SET hits = fizzbuzz_request.hits + 1").
		Exec(ctx)
	if err != nil {
		f.logger.Error("Failed to save FizzBuzzRequest", zap.Error(err))
		return errors.Wrap(err).WithKind("internal_error")
	}

	return nil
}

func (f *fizzBuzzRepository) GetMostHits(ctx context.Context) (domain.FizzbuzzRequest, errors.Error) {
	var fizzbuzzRequest domain.FizzbuzzRequest

	err := f.db.NewSelect().
		Model(&fizzbuzzRequest).
		Order("hits DESC").
		Limit(1).
		Scan(ctx)

	if err != nil {
		f.logger.Error("Failed to get most hits FizzBuzzRequest", zap.Error(err))
		return fizzbuzzRequest, errors.Wrap(err).WithKind("internal_error")
	}

	return fizzbuzzRequest, nil
}
