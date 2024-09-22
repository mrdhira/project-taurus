package redisExt

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedisExt interface {
	Ping(ctx context.Context) error
	Close() error

	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	HSet(ctx context.Context, key string, field string, values ...interface{}) error
	HGet(ctx context.Context, key string, field string) (string, error)
	HDel(ctx context.Context, key string, fields ...string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
}

type Config struct {
	Addr     string
	Password string
	DB       int
}

type redisExt struct {
	redisConn *redis.Client

	prefix string
}

func New(config Config) (IRedisExt, error) {
	var (
		redisConn *redis.Client
		err       error
	)
	redisConn = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	// Check connection
	_, err = redisConn.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}

	return &redisExt{
		redisConn: redisConn,
		prefix:    "taurus:",
	}, nil
}
