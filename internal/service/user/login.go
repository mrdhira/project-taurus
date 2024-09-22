package user

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/mrdhira/project-taurus/constant"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
	userModel "github.com/mrdhira/project-taurus/internal/model/user"
)

func (s *userService) Login(ctx context.Context, request *requestModel.UserLogin) (*responseModel.UserLogin, error) {
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

	// User not found
	if user == nil {
		s.logger.InfoContext(ctx, fmt.Sprintf("user with email %s not found", request.Email))
		return nil, errors.New(constant.ErrorUserNotFound)
	}

	// Check password
	if err = user.ComparePassword(request.Password); err != nil {
		s.logger.InfoContext(ctx, fmt.Sprintf("password for user with email %s is incorrect", request.Email))
		return nil, err
	}

	// Generate Access and Refresh Token
	// TODO: Make it configurable
	// Note: Redis always expired first if there is latency between generate token and save to redis
	accessTokenExpiry := 24 * time.Hour
	refreshTokenExpiry := 30 * 24 * time.Hour

	accessToken, accessTokenExpiredAt, refreshToken, refreshTokenExpiredAt, err := s.jwtExt.
		GenerateTokenLogin(ctx,
			user.ID.String(),
			user.Email,
			accessTokenExpiry,
			refreshTokenExpiry,
		)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to generate token", slog.String("error", err.Error()))
		return nil, err
	}

	// Save refresh token to redis
	// Encode key and token before saving to redis
	userIdEncode := base64.StdEncoding.EncodeToString([]byte(user.ID.String()))
	accessTokenEncode := base64.StdEncoding.EncodeToString([]byte(accessToken))
	refreshTokenEncode := base64.StdEncoding.EncodeToString([]byte(refreshToken))

	err = s.redisExt.Set(ctx, fmt.Sprintf("user:login:accessToken:%s", accessTokenEncode), userIdEncode, accessTokenExpiry)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to save access token to redis", slog.String("error", err.Error()))
		return nil, err
	}

	err = s.redisExt.Set(ctx, fmt.Sprintf("user:login:refreshToken:%s", refreshTokenEncode), userIdEncode, refreshTokenExpiry)
	if err != nil {
		// Delete access token if refresh token failed to save
		go s.redisExt.Del(ctx, userIdEncode)
		s.logger.ErrorContext(ctx, "failed to save refresh token to redis", slog.String("error", err.Error()))
		return nil, err
	}

	// TODO: Add activity log

	return &responseModel.UserLogin{
		AccessToken:      accessToken,
		AccessExpiredAt:  accessTokenExpiredAt,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshTokenExpiredAt,
	}, nil
}
