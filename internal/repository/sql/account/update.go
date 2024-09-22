package account

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mrdhira/project-taurus/constant"
	accountModel "github.com/mrdhira/project-taurus/internal/model/account"
)

var (
	updateQuery = `
		UPDATE accounts
		SET
			name = ?,
			status = ?,
			updated_at = NOW()
		WHERE id = ? AND user_id = ? AND deleted_at IS NULL
	`

	updateBalanceQuery = `
		UPDATE accounts
		SET 
			balance = ?,
			last_update_balance_at = NOW(),
			updated_at = NOW()
		WHERE id = ? AND user_id = ? AND deleted_at IS NULL
	`
)

func (r *accountRepository) Update(ctx context.Context, userID uuid.UUID, account *accountModel.Account) error {
	result, err := r.sqlExt.ExecContext(ctx, updateQuery, account.Name, account.Status, account.ID, userID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New(constant.ErrorAccountNotFound)
	}

	return nil
}

func (r *accountRepository) UpdateBalance(ctx context.Context, userID uuid.UUID, account *accountModel.Account) error {
	result, err := r.sqlExt.ExecContext(ctx, updateBalanceQuery, account.Balance, account.ID, userID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New(constant.ErrorAccountNotFound)
	}

	return nil
}
