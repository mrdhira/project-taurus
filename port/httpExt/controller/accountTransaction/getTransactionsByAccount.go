package accountTransaction

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/mrdhira/project-taurus/constant"
	requestModel "github.com/mrdhira/project-taurus/internal/model/request"
	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
)

func (c *accountTransactionController) GetTransactionsByAccount(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = r.Context()
		userID    = ctx.Value(constant.CtxKeyUserID).(string)
		accountID = r.PathValue("account_id")
		page      = r.URL.Query().Get("page")
		perPage   = r.URL.Query().Get("per_page")
		fromDate  = r.URL.Query().Get("from_date")
		toDate    = r.URL.Query().Get("to_date")
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

	pageParse, err := strconv.Atoi(page)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to parse page", slog.String("error", err.Error()), slog.String("page", page))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to parse page", err, nil)
		return
	}

	perPageParse, err := strconv.Atoi(perPage)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to parse per page", slog.String("error", err.Error()), slog.String("per_page", perPage))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to parse per page", err, nil)
		return
	}

	req := &requestModel.GetTransactionsByAccount{
		UserID:    userIDParse,
		AccountID: accountIDParse,
		Page:      pageParse,
		PerPage:   perPageParse,
	}

	if fromDate != "" {
		fromDateParse, err := time.Parse(time.RFC3339, fromDate)
		if err != nil {
			c.logger.ErrorContext(ctx, "failed to parse from date", slog.String("error", err.Error()), slog.String("from_date", fromDate))
			responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to parse from date", err, nil)
			return
		}
		req.FromDate = &fromDateParse
	}

	if toDate != "" {
		toDateParse, err := time.Parse(time.RFC3339, toDate)
		if err != nil {
			c.logger.ErrorContext(ctx, "failed to parse to date", slog.String("error", err.Error()), slog.String("to_date", toDate))
			responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to parse to date", err, nil)
			return
		}
		req.ToDate = &toDateParse
	}

	// Check if one of date is exists, both of it must be exists
	// If not, return error
	if (req.FromDate == nil && req.ToDate != nil) ||
		(req.FromDate != nil && req.ToDate == nil) {
		c.logger.ErrorContext(ctx, "failed to validate request", slog.String("error", "from_date and to_date must be exists at the same time"))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "from_date and to_date must be exists at the same time", err, nil)
		return
	}

	// Validate request
	err = c.validate.Struct(req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to validate request", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to validate request", err, nil)
		return
	}

	transactions, err := c.accountSvc.GetTransactionsByAccount(ctx, req)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to get transactions by account", slog.String("error", err.Error()))
		responseModel.NewHttpJSONResponse(w, http.StatusInternalServerError, "failed to get transactions by account", err, nil)
		return
	}

	responseModel.NewHttpJSONResponse(w, http.StatusOK, "success", nil, transactions)
}
