package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Redis is a Redis client wrapper.
type Redis struct {
	client *redis.Client
}

// Option is the Redis option.
type Option struct {
	Addr     string
	Password string
	DB       int
}

// New returns a new Redis client.
func New(opt *Option) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("redis ping error: %w", err)
	}
	return &Redis{client: client}, nil
}

// Client returns the Redis client.
func (r *Redis) Client() *redis.Client {
	return r.client
}

// Close closes the Redis client.
func (r *Redis) Close() error {
	return r.client.Close()
}
