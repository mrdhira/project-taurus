package healthCheck

import (
	"context"
	"time"

	responseModel "github.com/mrdhira/project-taurus/internal/model/response"
	"golang.org/x/sync/errgroup"
)

func (s *healthCheckService) HealthCheck(ctx context.Context) (*responseModel.HealthCheck, error) {
	var (
		response = &responseModel.HealthCheck{
			Status: "UP",
		}
		errG = new(errgroup.Group)
	)

	// DB health check
	errG.Go(func() error {
		response.DBStatus = "UP"
		response.DBCheckedAt = time.Now()
		err := s.sqlExt.Ping()
		if err != nil {
			response.DBStatus = "DOWN"
			return err
		}
		return nil
	})

	// Redis health check
	errG.Go(func() error {
		response.RedisStatus = "UP"
		response.RedisCheckedAt = time.Now()
		err := s.redisExt.Ping(ctx)
		if err != nil {
			response.RedisStatus = "DOWN"
			return err
		}
		return nil
	})

	err := errG.Wait()
	if err != nil {
		response.Status = "DOWN"
		return response, err
	}

	return response, nil
}
