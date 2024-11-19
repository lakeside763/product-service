package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lakeside763/product-service/config"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(addr string) *RedisCache {
	config := config.LoadConfig()

	client := redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})

	return &RedisCache{Client: client}
}

// Close method to close the Redis connection
func (c *RedisCache) Close() error {
	return c.Client.Close()
}

func (c *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return c.Client.Set(context.TODO(), key, value, expiration).Err()
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.Client.Get(context.TODO(), key).Result()
}