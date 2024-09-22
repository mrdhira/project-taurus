package request

import "github.com/google/uuid"

type CreateAccount struct {
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	AccountName string    `json:"account_name" validate:"required"`
}

type UpdateAccount struct {
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	AccountID   uuid.UUID `json:"account_id" validate:"required"`
	AccountName string    `json:"account_name" validate:"required"`
}

type DeleteAccount struct {
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	AccountID uuid.UUID `json:"account_id" validate:"required"`
}

type GetAccountByUser struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type GetAccountByID struct {
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	AccountID uuid.UUID `json:"account_id" validate:"required"`
}
