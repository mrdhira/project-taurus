package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mrdhira/project-taurus/constant"
)

var (
	deleteQuery = `
		UPDATE users
		SET 
			deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`
)

func (r *userRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	result, err := r.sqlExt.ExecContext(ctx, deleteQuery, userID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New(constant.ErrorUserNotFound)
	}

	return nil
}
