package accountTransaction

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/mrdhira/project-taurus/internal/service"
	"github.com/mrdhira/project-taurus/port/httpExt/controller"
)

type accountTransactionController struct {
	logger     *slog.Logger
	validate   *validator.Validate
	accountSvc service.IAccountTransactionService
}

type OptionFunc func(*accountTransactionController)

func New(
	logger *slog.Logger,
	validate *validator.Validate,
	opts ...OptionFunc,
) controller.IV1AccountTransactionController {
	ctrl := &accountTransactionController{
		logger:   logger,
		validate: validate,
	}

	for _, opt := range opts {
		opt(ctrl)
	}

	return ctrl
}

func WithAccountTransactionService(accountTransactionService service.IAccountTransactionService) OptionFunc {
	return func(c *accountTransactionController) {
		c.accountSvc = accountTransactionService
	}
}
