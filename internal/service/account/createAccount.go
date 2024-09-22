package account

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	accountModel "github.com/mrdhira/project-taurus/internal/model/account"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (s *accountService) CreateAccount(ctx context.Context, request *requestModel.CreateAccount) (*responseModel.Account, error) {
	account := &accountModel.Account{
		ID:      uuid.New(),
		UserID:  request.UserID,
		Name:    request.AccountName,
		Balance: decimal.Zero,
		Status:  accountModel.StatusActive,
	}

	err := s.accountRepo.Create(ctx, account)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to create account", slog.String("error", err.Error()))
		return nil, err
	}

	response := &responseModel.Account{
		AccountID:   account.ID.String(),
		AccountName: account.Name,
		CreatedAt:   account.CreatedAt,
	}

	return response, nil
}
