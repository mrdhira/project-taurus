package accountTransaction

import (
	"log/slog"

	"github.com/mrdhira/project-taurus/internal/repository"
	"github.com/mrdhira/project-taurus/internal/service"
	"github.com/mrdhira/project-taurus/pkg/redisExt"
)

type accountTransactionService struct {
	logger   *slog.Logger
	redisExt redisExt.IRedisExt

	accountRepo            repository.IAccountRepository
	accountTransactionRepo repository.IAccountTransactionRepository
	userRepo               repository.IUserRepository
}

type OptionFunc func(*accountTransactionService)

func New(
	logger *slog.Logger,
	opts ...OptionFunc,
) service.IAccountTransactionService {
	svc := &accountTransactionService{
		logger: logger,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func WithRedis(redisExt redisExt.IRedisExt) OptionFunc {
	return func(s *accountTransactionService) {
		s.redisExt = redisExt
	}
}

func WithAccountRepo(accountRepository repository.IAccountRepository) OptionFunc {
	return func(s *accountTransactionService) {
		s.accountRepo = accountRepository
	}
}

func WithAccountTransactionRepo(accountTransactionRepository repository.IAccountTransactionRepository) OptionFunc {
	return func(s *accountTransactionService) {
		s.accountTransactionRepo = accountTransactionRepository
	}
}

func WithUserRepo(userRepository repository.IUserRepository) OptionFunc {
	return func(s *accountTransactionService) {
		s.userRepo = userRepository
	}
}
