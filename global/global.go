package global

import (
	"embed"

	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/cache"
	"github.com/fitv/min/core/db"
	"github.com/fitv/min/core/logger"
	"github.com/fitv/min/ent"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
)

var App *app.Application

func FS() embed.FS {
	return App.FS
}

func Ent() *ent.Client {
	return App.DB.Client()
}

func DB() *db.DB {
	return App.DB
}

func Cache() cache.Cache {
	return App.Cache
}

func Redis() *redis.Client {
	return App.Redis.Client()
}

func Log() *logger.Logger {
	return App.Logger
}

func Trans() ut.Translator {
	return App.Translator
}
