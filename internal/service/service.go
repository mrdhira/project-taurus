package service

import (
	"context"

	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

type IHealthCheckService interface {
	HealthCheck(ctx context.Context) (*responseModel.HealthCheck, error)
}

type IAccountService interface {
	CreateAccount(ctx context.Context, request *requestModel.CreateAccount) (*responseModel.Account, error)
	GetAccountByID(ctx context.Context, request *requestModel.GetAccountByID) (*responseModel.AccountDetail, error)
	GetAccountsByUser(ctx context.Context, request *requestModel.GetAccountByUser) ([]*responseModel.Account, error)
	UpdateAccount(ctx context.Context, request *requestModel.UpdateAccount) (*responseModel.Account, error)
	DeleteAccount(ctx context.Context, request *requestModel.DeleteAccount) error
}

type IAccountTransactionService interface {
	CreateTransaction(ctx context.Context, request *requestModel.CreateTransaction) (*responseModel.Transaction, error)
	GetTransactionsByAccount(ctx context.Context, request *requestModel.GetTransactionsByAccount) ([]*responseModel.Transaction, error)
	UpdateTransaction(ctx context.Context, request *requestModel.UpdateTransaction) (*responseModel.Transaction, error)
	DeleteTransaction(ctx context.Context, request *requestModel.DeleteTransaction) error
}

type IUserService interface {
	Register(ctx context.Context, request *requestModel.UserRegister) (*responseModel.UserRegister, error)
	Login(ctx context.Context, request *requestModel.UserLogin) (*responseModel.UserLogin, error)
}
