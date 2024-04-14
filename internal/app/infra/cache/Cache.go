package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

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
	replicaHost := config.GetConfig("CACHE_HOST_REPLICAS")

	return &cacheConnection{
		client: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       0,
		}),
		replicasClient: redis.NewClient(&redis.Options{
			Addr:     replicaHost,
			Password: password,
			DB:       0,
		}),
		ttl: getDefaultTTLInSeconds(),
	}
}

func getDefaultTTLInSeconds() int {
	ttl := config.GetConfig("CACHE_TTL")
	ttlInt, err := strconv.Atoi(ttl)
	if err != nil {
		return 60
	}
	return ttlInt
}

type cacheConnection struct {
	client         *redis.Client
	replicasClient *redis.Client
	ttl            int // seconds
}

func (c *cacheConnection) Get(ctx context.Context, key string, value interface{}) error {
	result, err := c.replicasClient.Get(ctx, key).Result()

	if err != nil {
		return nil
	}

	return json.Unmarshal([]byte(result), value)
}

func (c *cacheConnection) Set(ctx context.Context, key string, value interface{}) error {
	statusCmd := c.client.Set(ctx, key, value, time.Duration(c.ttl)*time.Second)
	return statusCmd.Err()
}
