package account

import (
	"context"

	"github.com/google/uuid"
	accountModel "github.com/mrdhira/project-taurus/internal/model/account"
)

var (
	getByIDQuery = `
		SELECT
			id,
			user_id,
			name,
			balance,
			last_update_balance_at,
			status,
			created_at,
			updated_at
		FROM accounts
		WHERE
			id = ?
			user_id = ?
			AND deleted_at IS NULL
	`

	getByUserIDQuery = `
		SELECT
			id,
			user_id,
			name,
			status,
			created_at,
			updated_at
		FROM accounts
		WHERE 
			user_id = ?
			AND deleted_at IS NULL
	`
)

func (r *accountRepository) GetByID(ctx context.Context, userID uuid.UUID, accountID uuid.UUID) (*accountModel.Account, error) {
	var account accountModel.Account

	err := r.sqlExt.GetContext(ctx, &account, getByIDQuery, accountID, userID)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *accountRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*accountModel.Account, error) {
	accounts := make([]*accountModel.Account, 0)

	err := r.sqlExt.SelectContext(ctx, &accounts, getByUserIDQuery, userID)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
