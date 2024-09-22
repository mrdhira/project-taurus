package response

import (
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	AccountID   string    `json:"account_id"`
	AccountName string    `json:"account_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type AccountDetail struct {
	AccountID           string          `json:"account_id"`
	AccountName         string          `json:"account_name"`
	CurrentBalance      decimal.Decimal `json:"current_balance"`
	Balance             decimal.Decimal `json:"balance"`
	LastUpdateBalanceAt time.Time       `json:"last_update_balance_at"`
	CreatedAt           time.Time       `json:"created_at"`
}
