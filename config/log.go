package config

import (
	"github.com/fitv/min/core/config"
	"github.com/fitv/min/util/str"
)

type LogConfig struct {
	Path     string
	Filename string
	Level    string
	Daily    bool // Whether to generate a new log file every day
}

var Log = &LogConfig{
	Path:     config.GetString("log.path", "./logs"),
	Filename: config.GetString("log.filename", str.ToSnakeCase(App.Name)),
	Level:    config.GetString("log.level", "info"),
	Daily:    config.GetBool("log.daily", true),
}
