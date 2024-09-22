package middleware

import (
	"log/slog"
	"net/http"

	"github.com/mrdhira/project-taurus/pkg/jwtExt"
	"github.com/mrdhira/project-taurus/pkg/redisExt"
)

type IMiddleware interface {
	WithValidateAccessToken(next http.Handler) http.Handler
}

type OptionFunc func(*middleware)

type middleware struct {
	logger   *slog.Logger
	redisExt redisExt.IRedisExt
	jwtExt   jwtExt.IJwtExt
}

func New(
	logger *slog.Logger,
	opts ...OptionFunc,
) IMiddleware {
	m := &middleware{
		logger: logger,
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func WithRedis(redisExt redisExt.IRedisExt) OptionFunc {
	return func(m *middleware) {
		m.redisExt = redisExt
	}
}

func WithJwt(jwtExt jwtExt.IJwtExt) OptionFunc {
	return func(m *middleware) {
		m.jwtExt = jwtExt
	}
}
