package domain

import "github.com/mwm-io/gapi/errors"

type FizzBuzzInput struct {
	Int1  int    `json:"int1"  bun:"int1"`
	Int2  int    `json:"int2"  bun:"int2"`
	Limit int    `json:"limit" bun:"max_limit"`
	Str1  string `json:"str1"  bun:"str1"`
	Str2  string `json:"str2"  bun:"str2"`
}

func (f FizzBuzzInput) Validate() error {
	if f.Int1 == 0 {
		return errors.BadRequest("invalid_input", "int1 must be different than 0")
	}

	if f.Int2 == 0 {
		return errors.BadRequest("invalid_input", "int2 must be different than 0")
	}

	if f.Limit <= 0 {
		return errors.BadRequest("invalid_input", "limit must be greater than 0")
	}

	if f.Str1 == "" {
		return errors.BadRequest("invalid_input", "str1 must not be empty")
	}

	if f.Str2 == "" {
		return errors.BadRequest("invalid_input", "str2 must not be empty")
	}

	if f.Int1 == f.Int2 {
		return errors.BadRequest("invalid_input", "int1 and int2 must be different")
	}

	return nil
}

type FizzbuzzRequest struct {
	FizzBuzzInput
	Hits int `json:"hits" bun:"hits"`
}
