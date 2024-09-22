package account

import (
	"context"
	"log/slog"

	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
)

func (s *accountService) DeleteAccount(ctx context.Context, request *requestModel.DeleteAccount) error {
	err := s.accountRepo.Delete(ctx, request.AccountID, request.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to delete account", slog.String("error", err.Error()))
		return err
	}

	return nil
}
