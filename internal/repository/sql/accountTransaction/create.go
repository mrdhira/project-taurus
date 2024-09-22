package accountTransaction

import (
	"context"

	accountTransactionModel "github.com/mrdhira/project-taurus/internal/model/accountTransaction"
)

var (
	createQuery = `
		INSERT INTO account_transactions (
			id,
			account_id,
			name,
			amount,
			created_at,
			updated_at
		) VALUES (?,?,?,NOW(),NOW())
	`
)

func (r *accountTransactionRepository) Create(ctx context.Context, accountTransaction *accountTransactionModel.AccountTransaction) error {
	_, err := r.sqlExt.ExecContext(ctx, createQuery, accountTransaction.ID, accountTransaction.AccountID, accountTransaction.Name, accountTransaction.Amount)
	if err != nil {
		return err
	}

	return nil
}
