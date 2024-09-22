package response

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	TransactionID     string          `json:"transaction_id"`
	TransactionName   string          `json:"transaction_name"`
	TransactionAmount decimal.Decimal `json:"transaction_amount"`
	CreatedAt         time.Time       `json:"created_at"`
}
