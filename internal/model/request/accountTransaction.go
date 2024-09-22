package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateTransaction struct {
	UserID            uuid.UUID       `json:"user_id" validate:"required"`
	AccountID         uuid.UUID       `json:"account_id" validate:"required"`
	TransactionName   string          `json:"transaction_name" validate:"required"`
	TransactionAmount decimal.Decimal `json:"amount" validate:"required"`
}

type GetTransactionsByAccount struct {
	UserID    uuid.UUID  `json:"user_id" validate:"required"`
	AccountID uuid.UUID  `json:"account_id" validate:"required"`
	Page      int        `json:"page"`
	PerPage   int        `json:"per_page"`
	FromDate  *time.Time `json:"from_date"`
	ToDate    *time.Time `json:"to_date"`
}

type UpdateTransaction struct {
	UserID               uuid.UUID       `json:"user_id" validate:"required"`
	AccountID            uuid.UUID       `json:"account_id" validate:"required"`
	AccountTransactionID uuid.UUID       `json:"account_transaction_id" validate:"required"`
	TransactionName      string          `json:"transaction_name"`
	TransactionAmount    decimal.Decimal `json:"amount"`
}

type DeleteTransaction struct {
	UserID               uuid.UUID `json:"user_id" validate:"required"`
	AccountID            uuid.UUID `json:"account_id" validate:"required"`
	AccountTransactionID uuid.UUID `json:"account_transaction_id" validate:"required"`
}
