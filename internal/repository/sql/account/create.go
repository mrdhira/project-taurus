package account

import (
	"context"

	accountModel "github.com/mrdhira/project-taurus/internal/model/account"
)

var (
	createQuery = `
		INSERT INTO accounts (id, user_id, name, balance, last_update_balance_at, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, NOW(), ?, NOW(), NOW())
	`
)

func (r *accountRepository) Create(ctx context.Context, account *accountModel.Account) error {
	_, err := r.sqlExt.ExecContext(ctx, createQuery, account.ID, account.UserID, account.Name, account.Balance, account.Status)
	if err != nil {
		return err
	}

	return nil
}
