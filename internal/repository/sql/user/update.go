package user

import (
	"context"
	"errors"

	"github.com/mrdhira/project-taurus/constant"
	userModel "github.com/mrdhira/project-taurus/internal/model/user"
)

var (
	updateQuery = `
		UPDATE users
		SET
			full_name = ?,
			email = ?,
			status = ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	updatePasswordQuery = `
		UPDATE users
		SET
			password = ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	updatePINQuery = `
		UPDATE users
		SET
			pin = ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`
)

func (r *userRepository) Update(ctx context.Context, user *userModel.User) error {
	result, err := r.sqlExt.ExecContext(ctx, updateQuery, user.FullName, user.Email, user.Status, user.ID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New(constant.ErrorUserNotFound)
	}

	return nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, user *userModel.User) error {
	result, err := r.sqlExt.ExecContext(ctx, updatePasswordQuery, user.Password, user.ID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New(constant.ErrorUserNotFound)
	}

	return nil
}

func (r *userRepository) UpdatePIN(ctx context.Context, user *userModel.User) error {
	result, err := r.sqlExt.ExecContext(ctx, updatePINQuery, user.PIN, user.ID)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New(constant.ErrorUserNotFound)
	}

	return nil
}
