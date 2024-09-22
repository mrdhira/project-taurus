package account

import (
	"context"
	"log/slog"

	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (s *accountService) GetAccountByID(ctx context.Context, request *requestModel.GetAccountByID) (*responseModel.AccountDetail, error) {
	account, err := s.accountRepo.GetByID(ctx, request.AccountID, request.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get account by id", slog.String("error", err.Error()))
		return nil, err
	}

	totalCalcTrx, err := s.accountTransactionRepo.GetAllAfterTimestampByAccountID(ctx, account.ID, account.LastUpdateBalanceAt)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get total calculation transaction", slog.String("error", err.Error()))
		return nil, err
	}

	response := &responseModel.AccountDetail{
		AccountID:           account.ID.String(),
		AccountName:         account.Name,
		CurrentBalance:      account.Balance.Add(totalCalcTrx.TotalAmount),
		Balance:             account.Balance,
		LastUpdateBalanceAt: account.LastUpdateBalanceAt,
		CreatedAt:           account.CreatedAt,
	}

	return response, nil
}
