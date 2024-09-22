package redisExt

import (
	"context"
	"time"
)

func (r *redisExt) Ping(ctx context.Context) error {
	_, err := r.redisConn.Ping(ctx).Result()
	return err
}

func (r *redisExt) Close() error {
	return r.redisConn.Close()
}

func (r *redisExt) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	key = r.prefix + key
	return r.redisConn.Set(ctx, key, value, expiration).Err()
}

func (r *redisExt) Get(ctx context.Context, key string) (string, error) {
	key = r.prefix + key
	return r.redisConn.Get(ctx, key).Result()
}

func (r *redisExt) Del(ctx context.Context, key string) error {
	key = r.prefix + key
	return r.redisConn.Del(ctx, key).Err()
}

func (r *redisExt) HSet(ctx context.Context, key string, field string, values ...interface{}) error {
	key = r.prefix + key
	return r.redisConn.HSet(ctx, key, field, values).Err()
}

func (r *redisExt) HGet(ctx context.Context, key string, field string) (string, error) {
	key = r.prefix + key
	return r.redisConn.HGet(ctx, key, field).Result()
}

func (r *redisExt) HDel(ctx context.Context, key string, fields ...string) error {
	key = r.prefix + key
	return r.redisConn.HDel(ctx, key, fields...).Err()
}

func (r *redisExt) Expire(ctx context.Context, key string, expiration time.Duration) error {
	key = r.prefix + key
	return r.redisConn.Expire(ctx, key, expiration).Err()
}
