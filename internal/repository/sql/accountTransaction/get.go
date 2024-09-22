package accountTransaction

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/google/uuid"

	accountTransactionModel "github.com/mrdhira/project-taurus/internal/model/accountTransaction"
)

var (
	getByAccountIDQuery = `
		SELECT
			id,
			account_id,
			name,
			amount,
			created_at,
			updated_at,
			deleted_at
		FROM account_transactions
		WHERE 
			account_id = ?
			AND deleted_at IS NULL
	`

	getAllAfterTimestampByAccountIDQuery = `
		SELECT
			account_id,
			SUM(amount) as total_amount
			MAX(created_at) as last_transaction_at
			FROM account_transactions
			WHERE account_id = ?
				AND created_at > ?
				AND deleted_at IS NULL
			GROUP BY account_id
		`
)

func (r *accountTransactionRepository) GetByAccountID(ctx context.Context, accountID uuid.UUID, filter *accountTransactionModel.AccountTransactionFilter) ([]*accountTransactionModel.AccountTransaction, error) {
	accountTransactions := make([]*accountTransactionModel.AccountTransaction, 0)

	// Slice to hold query conditions
	var conditions []string
	var args []interface{}

	// Add conditions based on the filter
	if filter.FromCreatedAt != nil {
		conditions = append(conditions, "created_at >= ?")
		args = append(args, filter.FromCreatedAt)
	}

	if filter.ToCreatedAt != nil {
		conditions = append(conditions, "created_at <= ?")
		args = append(args, filter.ToCreatedAt)
	}

	// If there are conditions, append them to the query
	if len(conditions) > 0 {
		getByAccountIDQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Add limit and offset for pagination
	getByAccountIDQuery += " LIMIT ? OFFSET ?"

	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.Offset == 0 {
		filter.Offset = 0
	}

	args = append(args, filter.Limit, filter.Offset)

	// Log the query
	r.logger.Debug("get by account id query", slog.String("query", getByAccountIDQuery), slog.Any("args", args))

	err := r.sqlExt.SelectContext(ctx, &accountTransactions, getByAccountIDQuery, args...)
	if err != nil {
		return nil, err
	}

	return accountTransactions, nil
}

func (r *accountTransactionRepository) GetAllAfterTimestampByAccountID(ctx context.Context, accountID uuid.UUID, timestamp time.Time) (*accountTransactionModel.TotalCalculationAccountTransaction, error) {
	var totalCalcAccTransaction accountTransactionModel.TotalCalculationAccountTransaction

	err := r.sqlExt.GetContext(ctx, &totalCalcAccTransaction, getAllAfterTimestampByAccountIDQuery, accountID, timestamp)
	if err != nil {
		return nil, err
	}

	return &totalCalcAccTransaction, nil
}
