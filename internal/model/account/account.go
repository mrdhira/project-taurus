package account

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID                  uuid.UUID       `db:"id" json:"id"`
	UserID              uuid.UUID       `db:"user_id" json:"user_id"`
	Name                string          `db:"name" json:"name"`
	Balance             decimal.Decimal `db:"balance" json:"balance"`
	LastUpdateBalanceAt time.Time       `db:"last_update_balance_at" json:"last_update_balance_at"`
	Status              Status          `db:"status" json:"status"`
	CreatedAt           time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time       `db:"updated_at" json:"updated_at"`
	DeletedAt           *time.Time      `db:"deleted_at" json:"deleted_at"`
}
