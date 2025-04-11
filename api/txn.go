package api

import (
	"net/http"
	"time"

	"github.com/papaya147/expenses/db/sqlc"
)

func (s *Server) CreateTxn(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	input, err := Read[CreateTxnInput](w, r)
	if err != nil {
		return err
	}

	arg := sqlc.CreateTxnParams{
		Timestamp:   time.Unix(input.Timestamp, 0),
		Amount:      input.Amount,
		Category:    input.Category,
		Description: input.Description,
	}
	t, err := s.store.CreateTxn(ctx, arg)
	if err != nil {
		return err
	}

	Write(w, http.StatusOK, ToTxn(t))
	return nil
}

func (s *Server) ListTxns(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	input, err := Read[ListTxnsInput](w, r)
	if err != nil {
		return err
	}

	var startTime *time.Time
	if input.StartTime != nil {
		st := time.Unix(*input.StartTime, 0)
		startTime = &st
	}

	var endTime *time.Time
	if input.EndTime != nil {
		et := time.Unix(*input.EndTime, 0)
		endTime = &et
	}

	arg := sqlc.ListTxnsParams{
		StartTime: startTime,
		EndTime:   endTime,
		Category:  input.Category,
		Limit:     input.Limit,
		Offset:    input.Offset,
	}
	t, err := s.store.ListTxns(ctx, arg)
	if err != nil && sqlc.ErrorCode(err) != sqlc.NoDataFound {
		return err
	}

	res := make([]Txn, len(t))
	for i, txn := range t {
		res[i] = ToTxn(txn)
	}

	Write(w, http.StatusOK, res)
	return nil
}

func (s *Server) UpdateTxn(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	input, err := Read[UpdateTxnInput](w, r)
	if err != nil {
		return err
	}

	arg := sqlc.UpdateTxnParams{
		Timestamp:   time.Unix(input.Timestamp, 0),
		Amount:      input.Amount,
		Category:    input.Category,
		Description: input.Description,
		ID:          input.ID,
	}
	t, err := s.store.UpdateTxn(ctx, arg)
	if err != nil {
		return err
	}

	Write(w, http.StatusOK, ToTxn(t))
	return nil
}
