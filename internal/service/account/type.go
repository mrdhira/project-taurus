package account

import (
	"log/slog"

	"github.com/mrdhira/project-taurus/internal/repository"
	"github.com/mrdhira/project-taurus/internal/service"
	"github.com/mrdhira/project-taurus/pkg/redisExt"
)

type accountService struct {
	logger   *slog.Logger
	redisExt redisExt.IRedisExt

	accountRepo            repository.IAccountRepository
	accountTransactionRepo repository.IAccountTransactionRepository
	userRepo               repository.IUserRepository
}

type OptionFunc func(*accountService)

func New(
	logger *slog.Logger,
	opts ...OptionFunc,
) service.IAccountService {
	svc := &accountService{
		logger: logger,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func WithRedis(redisExt redisExt.IRedisExt) OptionFunc {
	return func(s *accountService) {
		s.redisExt = redisExt
	}
}

func WithAccountRepo(accountRepository repository.IAccountRepository) OptionFunc {
	return func(s *accountService) {
		s.accountRepo = accountRepository
	}
}

func WithAccountTransactionRepo(accountTransactionRepository repository.IAccountTransactionRepository) OptionFunc {
	return func(s *accountService) {
		s.accountTransactionRepo = accountTransactionRepository
	}
}

func WithUserRepo(userRepository repository.IUserRepository) OptionFunc {
	return func(s *accountService) {
		s.userRepo = userRepository
	}
}
