package accountTransaction

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/mrdhira/project-taurus/constant"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (c *accountTransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = r.Context()
		userID    = ctx.Value(constant.CtxKeyUserID).(string)
		accountID = r.PathValue("account_id")
	)

	// Decode request
	var req requestModel.CreateTransaction
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

	accountIDParse, err := uuid.Parse(accountID)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to parse account id", slog.String("error", err.Error()), slog.String("account_id", accountID))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to parse account id", err, nil)
		return
	}

	req.UserID = userIDParse
	req.AccountID = accountIDParse

	// Validate request
	err = c.validate.Struct(req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to validate request", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to validate request", err, nil)
		return
	}

	transaction, err := c.accountSvc.CreateTransaction(ctx, &req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to create transaction", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, constant.GetHTTPStatusCodeByError(err.Error()), "failed to create transaction", err, nil)
		return
	}

	responseModel.NewHttpJSONResponse(w, http.StatusOK, "transaction created successfully", nil, transaction)
}
