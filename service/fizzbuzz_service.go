package service

import (
	"context"
	"lbc/fizzbuzz/domain"
	"lbc/fizzbuzz/repository"
	"strconv"
	"strings"

	"github.com/mwm-io/gapi/errors"
)

type FizzBuzzService interface {
	GenerateFizzBuzz(input domain.FizzBuzzInput) (string, errors.Error)
}

type fizzBuzzService struct {
	fizzBuzzRepository repository.FizzBuzzRepository
}

func NewFizzBuzzService(fizzBuzzRepository repository.FizzBuzzRepository) FizzBuzzService {
	return &fizzBuzzService{
		fizzBuzzRepository: fizzBuzzRepository,
	}
}

func (f *fizzBuzzService) GenerateFizzBuzz(input domain.FizzBuzzInput) (string, errors.Error) {
	if err := input.Validate(); err != nil {
		return "", errors.Wrap(err).WithKind("invalid_input")
	}

	var result []string
	for i := 1; i <= input.Limit; i++ {
		switch {
		// Note: If Int1 and Int2 share factors, i%(Int1*Int2) == 0 won't work
		case i%input.Int1 == 0 && i%input.Int2 == 0:
			result = append(result, input.Str1+input.Str2)
		case i%input.Int1 == 0:
			result = append(result, input.Str1)
		case i%input.Int2 == 0:
			result = append(result, input.Str2)
		default:
			result = append(result, strconv.Itoa(i))
		}
	}

	if err := f.fizzBuzzRepository.Save(context.Background(), input); err != nil {
		return "", errors.Wrap(err).WithKind("internal_error")
	}

	return strings.Join(result, ","), nil
}
