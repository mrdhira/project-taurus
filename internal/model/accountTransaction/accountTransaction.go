package accountTransaction

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountTransaction struct {
	ID        uuid.UUID       `db:"id" json:"id"`
	AccountID uuid.UUID       `db:"account_id" json:"account_id"`
	Name      string          `db:"name" json:"name"`
	Amount    decimal.Decimal `db:"amount" json:"amount"`
	CreatedAt time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt time.Time       `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time      `db:"deleted_at" json:"deleted_at"`
}

type TotalCalculationAccountTransaction struct {
	AccountID         uuid.UUID       `db:"account_id" json:"account_id"`
	TotalAmount       decimal.Decimal `db:"total_amount" json:"total_amount"`
	LastTransactionAt time.Time       `db:"last_transaction_at" json:"last_transaction_at"`
}
