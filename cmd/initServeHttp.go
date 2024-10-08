package cmd

import (
	"log/slog"
	"os"

	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"

	"github.com/mrdhira/project-taurus/config"
	"github.com/mrdhira/project-taurus/pkg/jwtExt"
	"github.com/mrdhira/project-taurus/pkg/redisExt"
	"github.com/mrdhira/project-taurus/pkg/sqlExt"
	"github.com/mrdhira/project-taurus/pkg/sqlExt/mysql"
	"github.com/mrdhira/project-taurus/pkg/validatorExt"
)

func initServeHttpPackage(appConfig *config.AppConfig, appSecret *config.AppSecret) (
	*slog.Logger,
	sqlExt.ISqlExt,
	redisExt.IRedisExt,
	jwtExt.IJwtExt,
	*validator.Validate,
	error,
) {
	var (
		logger   *slog.Logger
		sql      sqlExt.ISqlExt
		redis    redisExt.IRedisExt
		jwt      jwtExt.IJwtExt
		validate *validator.Validate
		err      error
	)

	errG := new(errgroup.Group)

	// Logger
	errG.Go(func() error {
		// TODO: Will be standarized in the package module
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
		return nil
	})

	// MySQL
	errG.Go(func() error {
		sql, err = mysql.New(mysql.Config{
			Host:     appConfig.Database.Host,
			Port:     appConfig.Database.Port,
			Database: appConfig.Database.Database,
			Username: appSecret.Database.Username,
			Password: appSecret.Database.Password,
		})
		if err != nil {
			logger.Error("failed to connect to mysql", slog.String("error", err.Error()))
			return err
		}
		return nil
	})

	// Redis
	errG.Go(func() error {
		redis, err = redisExt.New(redisExt.Config{
			Addr:     appConfig.Redis.Addr,
			DB:       appConfig.Redis.DB,
			Password: appSecret.Redis.Password,
		})
		if err != nil {
			logger.Error("failed to connect to redis", slog.String("error", err.Error()))
			return err
		}
		return nil
	})

	// JWT
	errG.Go(func() error {
		jwt = jwtExt.New(
			appSecret.JWTSecretKey,
		)
		return nil
	})

	// Validator
	errG.Go(func() error {
		validate = validatorExt.New()
		return nil
	})

	if err := errG.Wait(); err != nil {
		if sql != nil {
			sql.Close()
		}

		if redis != nil {
			redis.Close()
		}

		return nil, nil, nil, nil, nil, err
	}

	return logger, sql, redis, jwt, validate, nil
}
