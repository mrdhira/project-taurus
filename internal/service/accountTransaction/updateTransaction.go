package accountTransaction

import (
	"context"
	"errors"
	"log/slog"

	"github.com/mrdhira/project-taurus/constant"
	accountTransactionModel "github.com/mrdhira/project-taurus/internal/model/accountTransaction"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (s *accountTransactionService) UpdateTransaction(ctx context.Context, request *requestModel.UpdateTransaction) (*responseModel.Transaction, error) {
	// Validate if account exists and user is the owner
	account, err := s.accountRepo.GetByID(ctx, request.UserID, request.AccountID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get account", slog.String("error", err.Error()))
		return nil, err
	}

	if account == nil {
		s.logger.ErrorContext(ctx, "account not found", slog.String("error", "account not found"))
		return nil, errors.New(constant.ErrorAccountNotFound)
	}

	// Update transaction
	transaction := &accountTransactionModel.AccountTransaction{
		ID:     request.AccountTransactionID,
		Name:   request.TransactionName,
		Amount: request.TransactionAmount,
	}

	err = s.accountTransactionRepo.Update(ctx, transaction)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to update transaction", slog.String("error", err.Error()))
		return nil, err
	}

	return &responseModel.Transaction{
		TransactionID:     transaction.ID.String(),
		TransactionName:   transaction.Name,
		TransactionAmount: transaction.Amount,
		CreatedAt:         transaction.CreatedAt,
	}, nil
}
