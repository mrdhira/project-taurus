package healthCheck

import (
	"log/slog"

	"github.com/mrdhira/project-taurus/internal/service"
	"github.com/mrdhira/project-taurus/pkg/redisExt"
	"github.com/mrdhira/project-taurus/pkg/sqlExt"
)

type healthCheckService struct {
	logger   *slog.Logger
	sqlExt   sqlExt.ISqlExt
	redisExt redisExt.IRedisExt
}

func New(
	logger *slog.Logger,
	sqlExt sqlExt.ISqlExt,
	redisExt redisExt.IRedisExt,
) service.IHealthCheckService {
	return &healthCheckService{
		logger:   logger,
		sqlExt:   sqlExt,
		redisExt: redisExt,
	}
}
