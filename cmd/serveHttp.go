package cmd

import (
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/mrdhira/project-taurus/config"
	"github.com/mrdhira/project-taurus/port/httpExt"
	"github.com/mrdhira/project-taurus/port/httpExt/middleware"

	// Repository
	accountRepository "github.com/mrdhira/project-taurus/internal/repository/sql/account"
	accountTransactionRepository "github.com/mrdhira/project-taurus/internal/repository/sql/accountTransaction"
	userRepository "github.com/mrdhira/project-taurus/internal/repository/sql/user"

	// Service
	accountService "github.com/mrdhira/project-taurus/internal/service/account"
	accountTransactionService "github.com/mrdhira/project-taurus/internal/service/accountTransaction"
	healthCheckService "github.com/mrdhira/project-taurus/internal/service/healthCheck"
	userService "github.com/mrdhira/project-taurus/internal/service/user"

	// Controller
	accountController "github.com/mrdhira/project-taurus/port/httpExt/controller/account"
	accountTransactionController "github.com/mrdhira/project-taurus/port/httpExt/controller/accountTransaction"
	healthCheckController "github.com/mrdhira/project-taurus/port/httpExt/controller/healthCheck"
	userController "github.com/mrdhira/project-taurus/port/httpExt/controller/user"
)

func init() {
	rootCmd.AddCommand(serveHttpCmd)
}

var serveHttpCmd = &cobra.Command{
	Use:   "serveHttp",
	Short: "Start HTTP server",
	Long:  `Start Taurus HTTP server`,
	Run: func(cmd *cobra.Command, args []string) {
		appConfig, appSecret, err := config.New(cfgFile, scrtFile)
		if err != nil {
			panic(err)
		}

		// Package
		logger, sql, redis, jwt, validate, err := initServeHttpPackage(appConfig, appSecret)
		if err != nil {
			panic(err)
		}
		defer sql.Close()
		defer redis.Close()

		// Repository
		accountRepo := accountRepository.New(logger, sql)
		accountTransactionRepo := accountTransactionRepository.New(logger, sql)
		userRepo := userRepository.New(logger, sql)

		// Service
		healthCheckSvc := healthCheckService.New(
			logger,
			sql,
			redis,
		)
		accountSvc := accountService.New(
			logger,
			accountService.WithRedis(redis),
			accountService.WithAccountRepo(accountRepo),
			accountService.WithAccountTransactionRepo(accountTransactionRepo),
			accountService.WithUserRepo(userRepo),
		)
		accountTransactionSvc := accountTransactionService.New(
			logger,
			accountTransactionService.WithRedis(redis),
			accountTransactionService.WithAccountRepo(accountRepo),
			accountTransactionService.WithAccountTransactionRepo(accountTransactionRepo),
			accountTransactionService.WithUserRepo(userRepo),
		)
		userSvc := userService.New(
			logger,
			userService.WithRedis(redis),
			userService.WithJwt(jwt),
			userService.WithUserRepo(userRepo),
		)

		// Middleware
		middleware := middleware.New(
			logger,
			middleware.WithRedis(redis),
			middleware.WithJwt(jwt),
		)

		// Controller
		healthCheckCtrl := healthCheckController.New(
			logger,
			healthCheckSvc,
		)
		accountCtrl := accountController.New(
			logger,
			validate,
			accountController.WithAccountService(accountSvc),
		)
		accountTransactionCtrl := accountTransactionController.New(
			logger,
			validate,
			accountTransactionController.WithAccountTransactionService(accountTransactionSvc),
		)
		userCtrl := userController.New(
			logger,
			validate,
			userController.WithUserService(userSvc),
		)

		mux := httpExt.New(
			middleware,
			healthCheckCtrl,
			accountCtrl,
			accountTransactionCtrl,
			userCtrl,
		)

		server := &http.Server{
			Addr:    ":8000",
			Handler: mux,
		}

		go func() {
			if err := server.ListenAndServe(); err != nil {
				logger.Error("failed to serve http", slog.String("error", err.Error()))
			}
		}()

		// Graceful shutdown
		quitSignal := make(chan os.Signal, 1)
		signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)

		<-quitSignal

		if err := server.Shutdown(nil); err != nil {
			logger.Error("failed to shutdown http server", slog.String("error", err.Error()))
		}

		logger.Info("server gracefully shutdown")
	},
}
