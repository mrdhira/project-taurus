package account

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mrdhira/project-taurus/constant"
)

var (
	deleteQuery = `
		UPDATE accounts
		SET
			deleted_at = NOW()
		WHERE id = ? AND user_id = ? AND deleted_at IS NULL
	`
)

func (r *accountRepository) Delete(ctx context.Context, userID uuid.UUID, accountID uuid.UUID) error {
	result, err := r.sqlExt.ExecContext(ctx, deleteQuery, accountID, userID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New(constant.ErrorAccountNotFound)
	}

	return nil
}
