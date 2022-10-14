package service

import (
	"fmt"

	"github.com/fitv/min/config"
	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/redis"
)

type Redis struct {
	Service
}

func (Redis) Register(app *app.Application) {
	var err error

	app.Redis, err = redis.New(&redis.Option{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
	if err != nil {
		panic(fmt.Errorf("redis init error: %w", err))
	}

	app.AddShutdown(func() {
		app.Redis.Close()
	})
}
