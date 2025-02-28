package cache

import (
	"context"
	"rbp/internal/config"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(cfg config.Redis) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &RedisClient{Client: client}
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(context.Background(), key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (int, error) {
	val, err := r.Client.Get(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}
