package account

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/mrdhira/project-taurus/internal/service"
	"github.com/mrdhira/project-taurus/port/httpExt/controller"
)

type accountController struct {
	logger     *slog.Logger
	validate   *validator.Validate
	accountSvc service.IAccountService
}

type OptionFunc func(*accountController)

func New(
	logger *slog.Logger,
	validate *validator.Validate,
	opts ...OptionFunc,
) controller.IV1AccountController {
	ctrl := &accountController{
		logger:   logger,
		validate: validate,
	}

	for _, opt := range opts {
		opt(ctrl)
	}

	return ctrl
}

func WithAccountService(accountService service.IAccountService) OptionFunc {
	return func(c *accountController) {
		c.accountSvc = accountService
	}
}
