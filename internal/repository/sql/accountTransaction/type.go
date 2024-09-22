package accountTransaction

import (
	"log/slog"

	"github.com/mrdhira/project-taurus/internal/repository"
	"github.com/mrdhira/project-taurus/pkg/sqlExt"
)

type accountTransactionRepository struct {
	logger *slog.Logger
	sqlExt sqlExt.ISqlExt
}

func New(
	logger *slog.Logger,
	sqlExt sqlExt.ISqlExt,
) repository.IAccountTransactionRepository {
	return &accountTransactionRepository{
		logger: logger,
		sqlExt: sqlExt,
	}
}
