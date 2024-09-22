package account

import (
	"context"
	"log/slog"

	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (s *accountService) UpdateAccount(ctx context.Context, request *requestModel.UpdateAccount) (*responseModel.Account, error) {
	account, err := s.accountRepo.GetByID(ctx, request.AccountID, request.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get account by id", slog.String("error", err.Error()))
		return nil, err
	}

	account.Name = request.AccountName

	err = s.accountRepo.Update(ctx, request.UserID, account)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to update account", slog.String("error", err.Error()))
		return nil, err
	}

	response := &responseModel.Account{
		AccountID:   account.ID.String(),
		AccountName: account.Name,
		CreatedAt:   account.CreatedAt,
	}

	return response, nil
}
