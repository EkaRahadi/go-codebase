package database

import (
	"context"
	"fmt"
	"time"

	"github.com/EkaRahadi/go-codebase/internal/config"
	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.Config) *redis.Client {
	redisCfg := cfg.Redis

	if cfg.App.Environment == constants.AppEnvironmentProduction {

		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", redisCfg.Host, redisCfg.Port),
			Password: redisCfg.Password,
		})

		return client
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisCfg.Host, redisCfg.Port),
		DB:       0,
		Password: redisCfg.Password,
	})

	return rdb
}

type RedisWrapper struct {
	redis *redis.Client
}

func NewRedisWrapper(rdb *redis.Client) *RedisWrapper {
	return &RedisWrapper{
		redis: rdb,
	}
}

func (rw *RedisWrapper) Get(ctx context.Context, key string) (string, error) {
	val, err := rw.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (rw *RedisWrapper) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	err := rw.redis.Set(ctx, key, value, ttl).Err()
	return err
}

func (rw *RedisWrapper) IncrementByOne(ctx context.Context, key string) error {
	err := rw.redis.Incr(ctx, key).Err()
	return err
}

func (rw *RedisWrapper) Delete(ctx context.Context, key string) error {
	result := rw.redis.Del(ctx, key)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
