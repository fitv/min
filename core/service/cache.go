package service

import (
	"fmt"

	"github.com/fitv/min/config"
	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/cache"
	"github.com/fitv/min/core/redis"
)

type Cache struct {
	Service
}

func (Cache) Register(app *app.Application) {
	switch config.Cache.Driver {
	case "redis":
		redis := redis.New(&redis.Option{
			Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
			Password: config.Redis.Password,
			DB:       config.Cache.Database,
		})
		app.AddClose(func() {
			redis.Close()
		})
		app.Cache = cache.NewRedisCache(redis.Client(), &cache.Option{
			Prefix: config.Cache.Prefix,
		})
	default:
		panic(fmt.Errorf("unsupported cache driver: %s", config.Cache.Driver))
	}
}
