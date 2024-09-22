package accountTransaction

import (
	"context"
	"errors"

	"github.com/mrdhira/project-taurus/constant"
	accountTransactionModel "github.com/mrdhira/project-taurus/internal/model/accountTransaction"
)

var (
	updateQuery = `
		UPDATE account_transactions
		SET
			name = ?,
			amount = ?,
			updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`
)

func (r *accountTransactionRepository) Update(ctx context.Context, accountTransaction *accountTransactionModel.AccountTransaction) error {
	result, err := r.sqlExt.ExecContext(ctx, updateQuery, accountTransaction.Name, accountTransaction.Amount, accountTransaction.ID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New(constant.ErrorTransactionNotFound)
	}

	return nil
}
