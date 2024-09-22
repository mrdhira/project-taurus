package user

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/mrdhira/project-taurus/constant"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
	userModel "github.com/mrdhira/project-taurus/internal/model/user"
)

func (s *userService) Register(ctx context.Context, request *requestModel.UserRegister) (*responseModel.UserRegister, error) {
	var (
		user *userModel.User
		err  error
	)

	// Check if email already exists
	user, err = s.userRepo.GetByEmail(ctx, request.Email)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get user by email", slog.String("error", err.Error()))
		return nil, err
	}

	// User already exists
	if user != nil {
		s.logger.InfoContext(ctx, fmt.Sprintf("user with email %s already exists", request.Email))
		return nil, errors.New(constant.ErrorUserAlreadyExists)
	}

	user = &userModel.User{
		ID:       uuid.New(),
		FullName: "",
		Email:    request.Email,
		Password: request.Password,
		Status:   userModel.StatusActive,
	}

	user.HashPassword()

	// Create user
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to create user", slog.String("error", err.Error()))
		return nil, err
	}

	return &responseModel.UserRegister{
		UserID:    user.ID.String(),
		FullName:  user.FullName,
		Email:     user.Email,
		Status:    user.StatusToString(),
		CreatedAt: user.CreatedAt,
	}, nil
}
