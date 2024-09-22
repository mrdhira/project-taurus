package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/mrdhira/project-taurus/constant"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (c *userController) Register(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	// Decode request
	var req requestModel.UserRegister
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to decode request", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to decode request", err, nil)
		return
	}

	// Validate request
	err = c.validate.Struct(req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to validate request", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to validate request", err, nil)
		return
	}

	resp, err := c.userSvc.Register(ctx, &req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to register user", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, constant.GetHTTPStatusCodeByError(err.Error()), "failed to register user", err, nil)
		return
	}

	responseModel.NewHttpJSONResponse(w, http.StatusOK, "user registered successfully", nil, resp)
}
