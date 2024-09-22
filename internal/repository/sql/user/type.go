package user

import (
	"log/slog"

	"github.com/mrdhira/project-taurus/internal/repository"
	"github.com/mrdhira/project-taurus/pkg/sqlExt"
)

type userRepository struct {
	logger *slog.Logger
	sqlExt sqlExt.ISqlExt
}

func New(
	logger *slog.Logger,
	sqlExt sqlExt.ISqlExt,
) repository.IUserRepository {
	return &userRepository{
		logger: logger,
		sqlExt: sqlExt,
	}
}
