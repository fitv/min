package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

type Option struct {
	Addr     string
	Password string
	DB       int
}

// New returns a new Redis client.
func New(opt *Option) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("redis connect error: %w", err))
	}
	return &Redis{client: client}
}

// Client returns the Redis client.
func (r *Redis) Client() *redis.Client {
	return r.client
}

// Close closes the Redis client.
func (r *Redis) Close() error {
	return r.client.Close()
}
