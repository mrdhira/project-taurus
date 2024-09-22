package account

import (
	"log/slog"

	"github.com/mrdhira/project-taurus/internal/repository"
	"github.com/mrdhira/project-taurus/pkg/sqlExt"
)

type accountRepository struct {
	logger *slog.Logger
	sqlExt sqlExt.ISqlExt
}

func New(
	logger *slog.Logger,
	sqlExt sqlExt.ISqlExt,
) repository.IAccountRepository {
	return &accountRepository{
		logger: logger,
		sqlExt: sqlExt,
	}
}
