package cache

import (
	"context"
	"time"
)

// Cache interface
type Cache interface {
	Get(ctx context.Context, key string) (string, bool)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration)
	Has(ctx context.Context, key string) bool
	Del(ctx context.Context, key string) bool
	TTL(ctx context.Context, key string) time.Duration
}

// Options for cache
type Option struct {
	Prefix string
}
