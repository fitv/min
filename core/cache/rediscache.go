package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ Cache = (*RedisCache)(nil)

// RedisCache is a Redis-based cache.
type RedisCache struct {
	client *redis.Client
	prefix string
}

// NewRedisCache creates a new RedisCache instance.
func NewRedisCache(client *redis.Client, opt *Option) *RedisCache {
	return &RedisCache{
		client: client,
		prefix: opt.Prefix,
	}
}

// Get returns the value for the given key in cache, panic when an error occurs.
func (r *RedisCache) Get(ctx context.Context, key string) (string, bool) {
	val, err := r.client.Get(ctx, r.realKey(key)).Result()
	if err == redis.Nil {
		return "", false
	}
	if err != nil {
		panic(fmt.Errorf("redis get error: %w", err))
	}
	return val, true
}

// Set sets the value for the given key into cache, panic when an error occurs.
func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) {
	err := r.client.SetEX(ctx, r.realKey(key), value, ttl).Err()
	if err != nil {
		panic(fmt.Errorf("redis set error: %w", err))
	}
}

// Has check if the cache key exists, panic when an error occurs.
func (r *RedisCache) Has(ctx context.Context, key string) bool {
	res, err := r.client.Exists(ctx, r.realKey(key)).Result()
	if err != nil {
		panic(fmt.Errorf("redis exists error: %w", err))
	}
	return res > 0
}

// TTL returns the remaining time to live of a key, panic when an error occurs.
func (r *RedisCache) TTL(ctx context.Context, key string) time.Duration {
	ttl, err := r.client.TTL(ctx, r.realKey(key)).Result()
	if err != nil {
		panic(fmt.Errorf("redis ttl error: %w", err))
	}
	if ttl < 0 {
		return 0
	}
	return ttl
}

// Del deletes the given key, panic when an error occurs.
func (r *RedisCache) Del(ctx context.Context, key string) bool {
	val, err := r.client.Del(ctx, r.realKey(key)).Result()
	if err != nil {
		panic(fmt.Errorf("redis del error: %w", err))
	}
	return val > 0
}

// realKey returns the key with prefix.
func (r *RedisCache) realKey(key string) string {
	return r.prefix + key
}
