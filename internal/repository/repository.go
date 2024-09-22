package repository

import (
	"context"
	"time"

	"github.com/google/uuid"

	accountModel "github.com/mrdhira/project-taurus/internal/model/account"
	accountTransactionModel "github.com/mrdhira/project-taurus/internal/model/accountTransaction"
	userModel "github.com/mrdhira/project-taurus/internal/model/user"
)

// Order by alphabetical order

type IAccountRepository interface {
	Create(ctx context.Context, account *accountModel.Account) error
	GetByID(ctx context.Context, userID uuid.UUID, accountID uuid.UUID) (*accountModel.Account, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*accountModel.Account, error)
	Update(ctx context.Context, userID uuid.UUID, account *accountModel.Account) error
	UpdateBalance(ctx context.Context, userID uuid.UUID, account *accountModel.Account) error
	Delete(ctx context.Context, userID uuid.UUID, accountID uuid.UUID) error
}

type IAccountTransactionRepository interface {
	Create(ctx context.Context, accountTransaction *accountTransactionModel.AccountTransaction) error
	GetByAccountID(ctx context.Context, accountID uuid.UUID, filter *accountTransactionModel.AccountTransactionFilter) ([]*accountTransactionModel.AccountTransaction, error)
	GetAllAfterTimestampByAccountID(ctx context.Context, accountID uuid.UUID, timestamp time.Time) (*accountTransactionModel.TotalCalculationAccountTransaction, error)
	Update(ctx context.Context, accountTransaction *accountTransactionModel.AccountTransaction) error
	Delete(ctx context.Context, accountTransactionID uuid.UUID) error
}

type IUserRepository interface {
	Create(ctx context.Context, user *userModel.User) error
	GetByID(ctx context.Context, userID uuid.UUID) (*userModel.User, error)
	GetByEmail(ctx context.Context, email string) (*userModel.User, error)
	Find(ctx context.Context, filter userModel.UserFilter) ([]*userModel.User, error)
	Update(ctx context.Context, user *userModel.User) error
	UpdatePassword(ctx context.Context, user *userModel.User) error
	UpdatePIN(ctx context.Context, user *userModel.User) error
	Delete(ctx context.Context, userID uuid.UUID) error
}
