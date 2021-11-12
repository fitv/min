package cache

import "time"

// Cache interface
type Cache interface {
	Get(key string) (string, bool)
	Set(key string, value interface{}, ttl time.Duration)
	Has(key string) bool
	Del(key string) bool
	TTL(key string) time.Duration
}

// Options for cache
type Option struct {
	Prefix string
}
