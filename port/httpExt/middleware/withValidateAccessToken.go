package middleware

import (
	"context"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/mrdhira/project-taurus/constant"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (m *middleware) WithValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
		)

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			m.logger.InfoContext(ctx, "unauthorized: missing token")
			responseModel.NewHttpJSONResponse(w, http.StatusUnauthorized, "unauthorized: missing token", nil, nil)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		tokenEncode := base64.StdEncoding.EncodeToString([]byte(token))

		userIdEncode, err := m.redisExt.Get(ctx, fmt.Sprintf("user:login:accessToken:%s", tokenEncode))
		if err != nil {
			m.logger.ErrorContext(ctx, "failed to get user id from redis", slog.String("error", err.Error()), slog.String("token", tokenEncode))
			responseModel.NewHttpJSONResponse(w, http.StatusInternalServerError, "failed to get user id from redis", err, nil)
			return
		}

		userId, err := base64.StdEncoding.DecodeString(userIdEncode)
		if err != nil {
			m.logger.ErrorContext(ctx, "failed to decode user id", slog.String("error", err.Error()), slog.String("user_id", string(userId)))
			responseModel.NewHttpJSONResponse(w, http.StatusInternalServerError, "failed to decode user id", err, nil)
			return
		}

		ctx = context.WithValue(ctx, constant.CtxKeyUserID, string(userId))

		accessTokenClaims, err := m.jwtExt.ValidateAccessToken(token)
		if err != nil {
			m.logger.ErrorContext(ctx, "failed to validate access token", slog.String("error", err.Error()))
			responseModel.NewHttpJSONResponse(w, http.StatusUnauthorized, "failed to validate access token", err, nil)
			return
		}

		ctx = context.WithValue(ctx, constant.CtxKeyEmail, accessTokenClaims.Email)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
