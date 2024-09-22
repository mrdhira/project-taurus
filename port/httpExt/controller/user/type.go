package user

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/mrdhira/project-taurus/internal/service"
	"github.com/mrdhira/project-taurus/port/httpExt/controller"
)

type userController struct {
	logger   *slog.Logger
	validate *validator.Validate
	userSvc  service.IUserService
}

type OptionFunc func(*userController)

func New(
	logger *slog.Logger,
	validate *validator.Validate,
	opts ...OptionFunc,
) controller.IV1UserController {
	ctrl := &userController{
		logger:   logger,
		validate: validate,
	}

	for _, opt := range opts {
		opt(ctrl)
	}

	return ctrl
}

func WithUserService(userService service.IUserService) OptionFunc {
	return func(c *userController) {
		c.userSvc = userService
	}
}
