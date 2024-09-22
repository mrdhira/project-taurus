package account

import (
	"context"
	"log/slog"

	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (s *accountService) GetAccountsByUser(ctx context.Context, request *requestModel.GetAccountByUser) ([]*responseModel.Account, error) {
	accounts, err := s.accountRepo.GetByUserID(ctx, request.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get account by user id", slog.String("error", err.Error()))
		return nil, err
	}

	response := make([]*responseModel.Account, 0)

	// TODO: Revamp and optimize this code
	// TODO: Check if need to use pagination
	for _, account := range accounts {
		responseAccount := &responseModel.Account{
			AccountID:   account.ID.String(),
			AccountName: account.Name,
			CreatedAt:   account.CreatedAt,
		}

		response = append(response, responseAccount)
	}

	return response, nil
}
