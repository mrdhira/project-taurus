package account

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/mrdhira/project-taurus/constant"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (c *accountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		userID = ctx.Value(constant.CtxKeyUserID).(string)
	)
	// Decode request
	var req requestModel.CreateAccount
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to decode request", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to decode request", err, nil)
		return
	}

	userIDParse, err := uuid.Parse(userID)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to parse user id", slog.String("error", err.Error()), slog.String("user_id", userID))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to parse user id", err, nil)
		return
	}

	req.UserID = userIDParse

	// Validate request
	err = c.validate.Struct(req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to validate request", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to validate request", err, nil)
		return
	}

	resp, err := c.accountSvc.CreateAccount(ctx, &req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to create account", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, constant.GetHTTPStatusCodeByError(err.Error()), "failed to create account", err, nil)
		return
	}

	responseModel.NewHttpJSONResponse(w, http.StatusOK, "account created successfully", nil, resp)
}
