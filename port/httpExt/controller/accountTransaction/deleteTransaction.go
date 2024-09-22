package accountTransaction

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/mrdhira/project-taurus/constant"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (c *accountTransactionController) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		ctx           = r.Context()
		userID        = ctx.Value(constant.CtxKeyUserID).(string)
		accountID     = r.PathValue("account_id")
		transactionID = r.PathValue("transaction_id")
	)

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

	transactionIDParse, err := uuid.Parse(transactionID)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to parse account transaction id", slog.String("error", err.Error()), slog.String("transaction_id", transactionID))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to parse account transaction  id", err, nil)
		return
	}

	req := &requestModel.DeleteTransaction{
		UserID:               userIDParse,
		AccountID:            accountIDParse,
		AccountTransactionID: transactionIDParse,
	}

	// Validate request
	err = c.validate.Struct(req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to validate request", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to validate request", err, nil)
		return
	}

	if err := c.accountSvc.DeleteTransaction(ctx, req); err != nil {
		c.logger.ErrorContext(ctx, "failed to delete transaction", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, constant.GetHTTPStatusCodeByError(err.Error()), "failed to delete transaction", err, nil)
		return
	}

	responseModel.NewHttpJSONResponse(w, http.StatusOK, "transaction deleted successfully", nil, nil)
}
