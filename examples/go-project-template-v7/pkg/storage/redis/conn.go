package redis

import (
	"context"
	"encoding/json"
	"time"

	"go-project-template-v5/config"

	"github.com/go-redis/redis/v8"
)

// connect

func NewRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
		DB:   cfg.Redis.Database,
	})

	return client
}

// utils

func SaveStructToRedis(ctx context.Context, client *redis.Client, key string, value interface{}, expiration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, jsonValue, expiration).Err()
}

func GetStructFromRedis(ctx context.Context, client *redis.Client, key string, dest interface{}) error {
	jsonValue, err := client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(jsonValue), dest)
}

func KeyExists(ctx context.Context, client *redis.Client, key string) (bool, error) {
	count, err := client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
