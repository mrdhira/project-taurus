package accountTransaction

import (
	"context"
	"errors"
	"log/slog"

	"github.com/mrdhira/project-taurus/constant"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
)

func (s *accountTransactionService) DeleteTransaction(ctx context.Context, request *requestModel.DeleteTransaction) error {
	// Validate if account exists and user is the owner
	account, err := s.accountRepo.GetByID(ctx, request.UserID, request.AccountID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get account", slog.String("error", err.Error()))
		return err
	}

	if account == nil {
		s.logger.ErrorContext(ctx, "account not found", slog.String("error", "account not found"))
		return errors.New(constant.ErrorAccountNotFound)
	}

	// Delete transaction
	err = s.accountTransactionRepo.Delete(ctx, request.AccountTransactionID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to delete transaction", slog.String("error", err.Error()))
		return err
	}

	return nil
}
