package sqlc

import (
	"database/sql"
	"errors"

	sqlite3 "github.com/mattn/go-sqlite3"
)

const (
	NoDataFound       = "no_data_found"
	ForeignKeyFailure = "foreign_key_violation"
	UniqueViolation   = "unique_violation"
)

func ErrorCode(err error) string {
	if errors.Is(err, sql.ErrNoRows) {
		return NoDataFound
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		switch sqliteErr.ExtendedCode {
		case sqlite3.ErrConstraintForeignKey:
			return ForeignKeyFailure
		case sqlite3.ErrConstraintUnique:
			return UniqueViolation
		default:
			return sqliteErr.Code.Error()
		}
	}

	return err.Error()
}
