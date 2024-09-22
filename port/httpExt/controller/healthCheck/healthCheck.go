package healthCheck

import (
	"log/slog"
	"net/http"

	"github.com/mrdhira/project-taurus/internal/model/response"
)

func (c *healthCheckController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp, err := c.healthCheckSvc.HealthCheck(r.Context())
	if err != nil {
		c.logger.Error("failed to health check", slog.String("error", err.Error()))
	}

	statusCode := http.StatusOK
	if resp.Status == "DOWN" {
		statusCode = http.StatusServiceUnavailable
	}

	response.NewHttpJSONResponse(w, statusCode, "OK", err, resp)
}
