package service

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fitv/go-logger"
	"github.com/fitv/min/config"
	"github.com/fitv/min/core/app"
	"github.com/fitv/min/util/file"
)

type Logger struct {
	Service
}

func (Logger) Register(app *app.Application) {
	err := file.MkdirAll(filepath.Dir(config.Log.Path))
	if err != nil {
		panic(fmt.Errorf("logger error: %w", err))
	}

	logLevel := logger.InfoLevel
	for key, val := range logger.LevelMap {
		if val == strings.ToLower(config.Log.Level) {
			logLevel = key
		}
	}
	option := &logger.Option{
		Path:  config.Log.Path,
		Daily: config.Log.Daily,
		Days:  config.Log.Days,
	}

	switch config.Log.Driver {
	case "file":
		app.Logger = logger.New(logger.NewFileLogger(option))
		app.Logger.SetLevel(logLevel)
	case "stdout":
		app.Logger = logger.New(logger.NewStdLogger())
		app.Logger.SetLevel(logLevel)
	default:
		panic(fmt.Errorf("logger driver %s not support", config.Log.Driver))
	}

	app.AddShutdown(func() {
		app.Logger.Close()
	})
}
