package user

import (
	"context"

	userModel "github.com/mrdhira/project-taurus/internal/model/user"
)

var (
	createQuery = `
		INSERT INTO users (id, full_name, email, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, NOW(), NOW())
	`
)

func (r *userRepository) Create(ctx context.Context, user *userModel.User) error {
	_, err := r.sqlExt.ExecContext(ctx, createQuery, user.ID, user.FullName, user.Email, user.Status)
	if err != nil {
		return err
	}

	return nil
}
