package api

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/papaya147/expenses/db/sqlc"
)

type Server struct {
	store     *sqlc.Queries
	router    chi.Router
	templates *template.Template
}

func NewServer(store *sqlc.Queries) *Server {
	s := &Server{store: store}

	funcMap := template.FuncMap{
		"divideBy": func(a int64, b float64) float64 {
			return float64(a) / b
		},
		"now": func() string {
			return time.Now().Format("2006-01-02T15:04")
		},
		"expenseOrIncome": func(a int64) string {
			if a < 0 {
				return "expense"
			}
			return "income"
		},
	}

	var err error
	s.templates, err = template.New("").Funcs(funcMap).ParseGlob("templates/*.html")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/txn", wrap(s.CreateTxn))
	r.Put("/txn", wrap(s.UpdateTxn))
	r.Post("/txn/list", wrap(s.ListTxns))
	r.Get("/txn/chart", wrap(s.TxnChart))

	r.Post("/category", wrap(s.CreateCategory))
	r.Post("/category/list", wrap(s.ListCategories))

	r.Get("/", wrap(s.handleIndex))

	fileServer := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	s.router = r
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
