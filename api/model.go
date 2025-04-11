package api

import (
	"time"
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
	Timestamp   string  `schema:"timestamp"`
	AmountPaisa float64 `schema:"amount"`
	Category    string  `schema:"category"`
	Description *string `schema:"description"`
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
	ID          int64     `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Amount      int64     `json:"amount"`
	Category    string    `json:"category"`
	Description *string   `json:"description"`
}

type TxnChart struct {
	Timestamps []string  `json:"timestamps"`
	Amounts    []float64 `json:"amounts"`
}
