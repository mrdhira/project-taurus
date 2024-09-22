package httpExt

import (
	"net/http"

	"github.com/mrdhira/project-taurus/port/httpExt/controller"
	"github.com/mrdhira/project-taurus/port/httpExt/middleware"
	"github.com/mrdhira/project-taurus/port/httpExt/router"
)

func New(
	middleware middleware.IMiddleware,
	healthCheckCtrl controller.IHealthCheckController,
	accountCtrl controller.IV1AccountController,
	accountTransactionCtrl controller.IV1AccountTransactionController,
	userCtrl controller.IV1UserController,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health-check", healthCheckCtrl.HealthCheck)

	mux.Handle("/api/v1", router.V1Group(
		middleware,
		accountCtrl,
		accountTransactionCtrl,
		userCtrl,
	))

	return mux
}
