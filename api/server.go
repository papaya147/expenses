package api

import "github.com/papaya147/expenses/db/sqlc"

type Server struct {
	store *sqlc.Queries
}
