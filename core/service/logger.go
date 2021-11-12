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

	fileLogger := logger.NewFileLogger(&logger.Option{
		Filename: config.Log.Filename,
		Path:     config.Log.Path,
		Daily:    config.Log.Daily,
	})
	logLevel := logger.InfoLevel

	for key, val := range logger.LevelMap {
		if val == strings.ToUpper(config.Log.Level) {
			logLevel = key
		}
	}
	app.Logger = logger.New(logLevel, fileLogger)

	app.AddClose(func() {
		app.Logger.Close()
	})
}
