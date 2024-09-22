package user

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	userModel "github.com/mrdhira/project-taurus/internal/model/user"
)

var (
	getByIDQuery = `
		SELECT
			id,
			full_name,
			email,
			password,
			pin,
			status,
			created_at,
			updated_at
		FROM
			users
		WHERE
			deleted_at IS NULL
			AND id = ?
	`

	getByEmailQuery = `
		SELECT
			id,
			full_name,
			email,
			password,
			pin,
			status,
			created_at,
			updated_at
		FROM
			users
		WHERE
			deleted_at IS NULL
			AND email = ?
	`

	findQuery = `
		SELECT
			id,
			full_name,
			email,
			password,
			pin,
			status,
			created_at,
			updated_at
		FROM
			users
		WHERE
			delete_at IS NULL
	`
)

func (r *userRepository) GetByID(ctx context.Context, userID uuid.UUID) (*userModel.User, error) {
	var user userModel.User

	err := r.sqlExt.GetContext(ctx, &user, getByIDQuery, userID.String())
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*userModel.User, error) {
	var user userModel.User

	err := r.sqlExt.GetContext(ctx, &user, getByEmailQuery, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Find(ctx context.Context, filter userModel.UserFilter) ([]*userModel.User, error) {
	users := make([]*userModel.User, 0)

	// Slice to hold query conditions
	var conditions []string
	var args []interface{}

	// Add conditions based on the filter
	if filter.ID != "" {
		conditions = append(conditions, "id LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", filter.ID)) // LIKE '%value%'
	}

	if filter.FullName != "" {
		conditions = append(conditions, "full_name LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", filter.FullName)) // LIKE '%value%'
	}

	if filter.Email != "" {
		conditions = append(conditions, "email LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", filter.Email)) // LIKE '%value%'
	}

	if filter.Status != "" {
		conditions = append(conditions, "status = ?")
		args = append(args, filter.Status) // Exact match
	}

	// If there are conditions, append them to the query
	if len(conditions) > 0 {
		findQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Add limit and offset for pagination
	findQuery += " LIMIT ? OFFSET ?"

	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.Offset == 0 {
		filter.Offset = 0
	}

	args = append(args, filter.Limit, filter.Offset)

	// Log the query
	r.logger.Debug("find query", slog.String("query", findQuery), slog.Any("args", args))

	err := r.sqlExt.SelectContext(ctx, &users, findQuery, args...)
	if err != nil {
		return nil, err
	}

	return users, nil
}
