package db

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ForeignKeyViolation = "foreign_key_violation"
	UniqueViolation     = "unique_violation"
)

var ErrRecordNotFound = pgx.ErrNoRows

var ErrUniqueViolation = &pgconn.PgError{
	Code: UniqueViolation,
}

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}
