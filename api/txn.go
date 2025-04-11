package api

import (
	"math"
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

	timestamp, err := time.Parse("2006-01-02T15:04", input.Timestamp)
	if err != nil {
		return err
	}

	amount := int64(math.Round(input.AmountPaisa * 100))

	arg := sqlc.CreateTxnParams{
		Timestamp:   timestamp,
		Amount:      amount,
		Category:    input.Category,
		Description: input.Description,
	}
	t, err := s.store.CreateTxn(ctx, arg)
	if err != nil {
		return err
	}

	if err := s.templates.ExecuteTemplate(w, "expense-item", Txn(t)); err != nil {
		return err
	}

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
		res[i] = Txn(txn)
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

	if err := s.templates.ExecuteTemplate(w, "expense-item", Txn(t)); err != nil {
		return err
	}

	return nil
}

func (s *Server) TxnChart(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	ts, err := s.store.ListTxnsByDate(ctx)
	if err != nil && sqlc.ErrorCode(err) != sqlc.NoDataFound {
		return err
	}

	timestamps := make([]string, len(ts))
	amounts := make([]float64, len(ts))
	for i, t := range ts {
		timestamps[i] = t.Date.(string)
		amounts[i] = float64(t.Amount) / 100
	}

	Write(w, http.StatusOK, TxnChart{
		Timestamps: timestamps,
		Amounts:    amounts,
	})
	return nil
}
