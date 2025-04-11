package api

import (
	"net/http"

	"github.com/papaya147/expenses/db/sqlc"
)

func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	input, err := Read[CreateCategoryInput](w, r)
	if err != nil {
		return err
	}

	c, err := s.store.CreateCategory(ctx, input.Name)
	if err != nil {
		return err
	}

	Write(w, http.StatusOK, Category{Name: c})
	return nil
}

func (s *Server) ListCategories(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	input, err := Read[ListCategoriesInput](w, r)
	if err != nil {
		return err
	}

	arg := sqlc.ListCategoriesParams{
		Name:  input.Name,
		Limit: 100,
	}
	cs, err := s.store.ListCategories(ctx, arg)
	if err != nil && sqlc.ErrorCode(err) != sqlc.NoDataFound {
		return err
	}

	res := make([]Category, len(cs))
	for i, c := range cs {
		res[i] = Category{Name: c}
	}

	Write(w, http.StatusOK, res)
	return nil
}
