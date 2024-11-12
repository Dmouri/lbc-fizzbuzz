package utils

import (
	"context"
	"os"
	"path/filepath"

	"github.com/mwm-io/gapi/errors"
	"github.com/uptrace/bun"
)

func LoadFixtures(db *bun.DB) errors.Error {
	path, err := filepath.Abs("../testdata/fixtures/fizzbuzz_hits.sql")
	if err != nil {
		return errors.Wrap(err).WithKind("path_error")
	}

	fixtures, err := os.ReadFile(path)
	if err != nil {
		return errors.Wrap(err).WithKind("read_error")
	}
	_, err = db.ExecContext(context.Background(), string(fixtures))
	if err != nil {
		return errors.Wrap(err).WithKind("exec_error")
	}

	return nil
}

func ResetDatabase(db *bun.DB) errors.Error {
	_, err := db.Exec("TRUNCATE TABLE fizzbuzz_requests RESTART IDENTITY")
	if err != nil {
		return errors.Wrap(err).WithKind("truncate_error")
	}

	return nil
}
