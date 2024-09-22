package accountTransaction

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/mrdhira/project-taurus/constant"
	accountTransactionModel "github.com/mrdhira/project-taurus/internal/model/accountTransaction"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (s *accountTransactionService) GetTransactionsByAccount(ctx context.Context, request *requestModel.GetTransactionsByAccount) ([]*responseModel.Transaction, error) {
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

	// If page is less than 1, set default to 1
	if request.Page < 1 {
		request.Page = 1
	}

	// If per page is less than 1, set default to 10
	if request.PerPage < 1 {
		request.PerPage = 10
	}

	filter := &accountTransactionModel.AccountTransactionFilter{
		Offset: (request.Page - 1) * request.PerPage,
		Limit:  request.PerPage,
	}

	// If from date is nil, set default to 30 days ago
	if request.FromDate == nil {
		fromDate := time.Now().AddDate(0, 0, -30)
		filter.FromCreatedAt = &fromDate
	} else {
		filter.FromCreatedAt = request.FromDate
	}

	// If to date is nil, set default to now
	if request.ToDate == nil {
		toDate := time.Now()
		filter.ToCreatedAt = &toDate
	} else {
		filter.ToCreatedAt = request.ToDate
	}

	transactions, err := s.accountTransactionRepo.GetByAccountID(ctx, request.AccountID, filter)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get transactions by account id", slog.String("error", err.Error()))
		return nil, err
	}

	response := make([]*responseModel.Transaction, 0)

	// TODO: Revamp and optimize this code
	for _, transaction := range transactions {
		responseTransaction := &responseModel.Transaction{
			TransactionID:     transaction.ID.String(),
			TransactionName:   transaction.Name,
			TransactionAmount: transaction.Amount,
			CreatedAt:         transaction.CreatedAt,
		}

		response = append(response, responseTransaction)
	}

	return response, nil
}
