package jwtExt

import (
	"context"
	"time"
)

type IJwtExt interface {
	GenerateTokenLogin(
		ctx context.Context,
		userID string,
		email string,
		accessTokenExpiry time.Duration,
		refreshTokenExpiry time.Duration,
	) (string, time.Time, string, time.Time, error)
	ValidateAccessToken(tokenString string) (*accessTokenClaims, error)
}

type jwtExt struct {
	secretKey string
}

type OptionFunc func(*jwtExt)

func New(
	secretKey string,
	opts ...OptionFunc,
) IJwtExt {
	return &jwtExt{
		secretKey: secretKey,
	}
}
