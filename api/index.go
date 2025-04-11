package api

import (
	"net/http"

	"github.com/papaya147/expenses/db/sqlc"
)

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	arg := sqlc.ListTxnsParams{
		Limit:  1000,
		Offset: 0,
	}
	ts, err := s.store.ListTxns(ctx, arg)
	if err != nil && sqlc.ErrorCode(err) != sqlc.NoDataFound {
		return err
	}

	res := make([]Txn, len(ts))
	for i, t := range ts {
		res[i] = Txn(t)
	}

	data := map[string]any{
		"Expenses": res,
	}

	err = s.templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		return err
	}

	return nil
}
