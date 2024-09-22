package jwtExt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type accessTokenClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type refreshTokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *jwtExt) GenerateTokenLogin(
	ctx context.Context,
	userID string,
	email string,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,
) (string, time.Time, string, time.Time, error) {
	// Set default access token expiry to 24 hours if not set
	if accessTokenExpiry == 0 {
		accessTokenExpiry = 24 * time.Hour
	}

	accessTokenExpiredAt := time.Now().Add(accessTokenExpiry)

	// Generate Access Token
	accessClaims := accessTokenClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiredAt),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(j.secretKey)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, err
	}

	// Set default refresh token expiry to 30 days if not set
	if refreshTokenExpiry == 0 {
		refreshTokenExpiry = 30 * 24 * time.Hour
	}

	refreshTokenExpiredAt := time.Now().Add(refreshTokenExpiry)

	// Generate Refresh token
	refreshClaims := refreshTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpiredAt),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(j.secretKey)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, err
	}

	return accessTokenString, accessTokenExpiredAt, refreshTokenString, refreshTokenExpiredAt, nil
}

func (j *jwtExt) ValidateAccessToken(tokenString string) (*accessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &accessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*accessTokenClaims); ok && token.Valid {
		// Check if token is expired
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, jwt.ErrTokenExpired
		}

		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
