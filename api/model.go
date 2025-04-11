package api

import (
	"time"

	"github.com/papaya147/expenses/db/sqlc"
)

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type ListCategoriesInput struct {
	Name *string `json:"name"`
}

type Category struct {
	Name string `json:"name"`
}

type CreateTxnInput struct {
	Timestamp   int64   `json:"timestamp"`
	Amount      int64   `json:"amount"`
	Category    string  `json:"category"`
	Description *string `json:"description"`
}

type ListTxnsInput struct {
	StartTime *int64  `json:"start_time"`
	EndTime   *int64  `json:"end_time"`
	Category  *string `json:"category"`
	Limit     int64   `json:"limit"`
	Offset    int64   `json:"offset"`
}

type UpdateTxnInput struct {
	ID          int64   `json:"id"`
	Timestamp   int64   `json:"timestamp"`
	Amount      int64   `json:"amount"`
	Category    string  `json:"category"`
	Description *string `json:"description"`
}

type Txn struct {
	ID          int64   `json:"id"`
	Timestamp   string  `json:"timestamp"`
	Amount      int64   `json:"amount"`
	Category    string  `json:"category"`
	Description *string `json:"description"`
}

func ToTxn(txn sqlc.Txn) Txn {
	return Txn{
		ID:          txn.ID,
		Timestamp:   txn.Timestamp.Format(time.DateTime),
		Amount:      txn.Amount,
		Category:    txn.Category,
		Description: txn.Description,
	}
}
