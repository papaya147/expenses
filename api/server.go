package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/papaya147/expenses/db/sqlc"
)

type Server struct {
	store  *sqlc.Queries
	router chi.Router
}

func NewServer(store *sqlc.Queries) *Server {
	s := &Server{store: store}

	mux := chi.NewRouter()

	mux.Post("/txn", wrap(s.CreateTxn))
	mux.Put("/txn", wrap(s.UpdateTxn))
	mux.Post("/txn/list", wrap(s.ListTxns))

	mux.Post("/category", wrap(s.CreateCategory))
	mux.Post("/category/list", wrap(s.ListCategories))

	s.router = mux

	return s
}

func (s *Server) Start(address string) error {
	return http.ListenAndServe(address, s.router)
}

func wrap(fn func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			fmt.Println(err)
			Error(w, err)
		}
	}
}
