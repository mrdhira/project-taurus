package user

import (
	"log/slog"

	"github.com/mrdhira/project-taurus/internal/repository"
	"github.com/mrdhira/project-taurus/internal/service"
	"github.com/mrdhira/project-taurus/pkg/jwtExt"
	"github.com/mrdhira/project-taurus/pkg/redisExt"
)

type userService struct {
	logger   *slog.Logger
	redisExt redisExt.IRedisExt
	jwtExt   jwtExt.IJwtExt
	userRepo repository.IUserRepository
}

type OptionFunc func(*userService)

func New(
	logger *slog.Logger,
	opts ...OptionFunc,
) service.IUserService {
	svc := &userService{
		logger: logger,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func WithRedis(redisExt redisExt.IRedisExt) OptionFunc {
	return func(s *userService) {
		s.redisExt = redisExt
	}
}

func WithJwt(jwtExt jwtExt.IJwtExt) OptionFunc {
	return func(s *userService) {
		s.jwtExt = jwtExt
	}
}

func WithUserRepo(userRepository repository.IUserRepository) OptionFunc {
	return func(s *userService) {
		s.userRepo = userRepository
	}
}
