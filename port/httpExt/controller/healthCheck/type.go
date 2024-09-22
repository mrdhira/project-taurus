package healthCheck

import (
	"log/slog"

	"github.com/mrdhira/project-taurus/internal/service"
	"github.com/mrdhira/project-taurus/port/httpExt/controller"
)

type healthCheckController struct {
	logger         *slog.Logger
	healthCheckSvc service.IHealthCheckService
}

func New(
	logger *slog.Logger,
	healthCheckService service.IHealthCheckService,
) controller.IHealthCheckController {
	return &healthCheckController{
		logger:         logger,
		healthCheckSvc: healthCheckService,
	}
}
