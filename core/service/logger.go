package service

import (
	"fmt"
	"strings"

	"github.com/fitv/min/config"
	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/logger"
	"github.com/fitv/min/util/file"
)

type Logger struct {
	Service
}

func (Logger) Register(app *app.Application) {
	err := file.MkdirAll(config.Log.Path)
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
		Filename: config.Log.Filename,
		Path:     config.Log.Path,
		Daily:    config.Log.Daily,
	}

	switch config.Log.Driver {
	case "file":
		app.Logger = logger.New(logLevel, logger.NewFileLogger(option))
	case "zap":
		zapLogger, err := logger.NewZapLogger(option)
		if err != nil {
			panic(fmt.Errorf("zap logger error: %w", err))
		}
		app.Logger = logger.New(logLevel, zapLogger)
	default:
		panic(fmt.Errorf("logger driver %s not support", config.Log.Driver))
	}

	app.AddClose(func() {
		app.Logger.Close()
	})
}
