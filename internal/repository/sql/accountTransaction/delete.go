package accountTransaction

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mrdhira/project-taurus/constant"
)

var (
	deleteQuery = `
		UPDATE account_transactions
			SET 
				deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`
)

func (r *accountTransactionRepository) Delete(ctx context.Context, accountTransactionID uuid.UUID) error {
	result, err := r.sqlExt.ExecContext(ctx, deleteQuery, accountTransactionID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New(constant.ErrorTransactionNotFound)
	}

	return nil
}
