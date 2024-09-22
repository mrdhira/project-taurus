package account

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/mrdhira/project-taurus/constant"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (c *accountController) GetAccountsByUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		userID = ctx.Value(constant.CtxKeyUserID).(string)
	)

	userIDParse, err := uuid.Parse(userID)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to parse user id", slog.String("error", err.Error()), slog.String("user_id", userID))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to parse user id", err, nil)
		return
	}

	req := &requestModel.GetAccountByUser{
		UserID: userIDParse,
	}

	// Validate request
	err = c.validate.Struct(req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to validate request", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to validate request", err, nil)
		return
	}

	accounts, err := c.accountSvc.GetAccountsByUser(ctx, req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to get accounts by user", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, constant.GetHTTPStatusCodeByError(err.Error()), "failed to get accounts by user", err, nil)
		return
	}

	responseModel.NewHttpJSONResponse(w, http.StatusOK, "get accounts by user successfully", nil, accounts)
}
