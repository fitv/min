package service

import (
	"fmt"
	"strings"

	"github.com/fitv/go-logger"
	"github.com/fitv/min/config"
	"github.com/fitv/min/core/app"
)

type Logger struct {
	Service
}

func (Logger) Register(app *app.Application) {
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
		fileWriter := logger.NewFileWriter(option)
		app.Logger = logger.New()
		app.Logger.SetOutput(fileWriter)
		app.Logger.SetLevel(logLevel)

		app.AddShutdown(func() {
			fileWriter.Close()
		})
	default:
		panic(fmt.Errorf("logger driver %s not support", config.Log.Driver))
	}
}
