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
	app.Redis = redis.New(&redis.Option{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})

	app.AddClose(func() {
		app.Redis.Close()
	})
}
