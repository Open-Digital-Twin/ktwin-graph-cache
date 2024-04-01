package cache

import (
	"context"
	"encoding/json"

	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/config"
	redis "github.com/redis/go-redis/v9"
)

type CacheConnection interface {
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
}

func NewCacheConnection() CacheConnection {
	host := config.GetConfig("CACHE_HOST")
	password := config.GetConfig("CACHE_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return &cacheConnection{
		client: client,
	}
}

type cacheConnection struct {
	client *redis.Client
}

func (c *cacheConnection) Get(ctx context.Context, key string, value interface{}) error {
	result, err := c.client.Get(ctx, key).Result()

	if err != nil {
		return nil
	}

	return json.Unmarshal([]byte(result), value)
}

func (c *cacheConnection) Set(ctx context.Context, key string, value interface{}) error {
	statusCmd := c.client.Set(ctx, key, value, 0)
	return statusCmd.Err()
}
