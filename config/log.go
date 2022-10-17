package config

import (
	"github.com/fitv/min/core/config"
	"github.com/fitv/min/util/str"
)

type LogConfig struct {
	Driver string // Supports: "file", "stdout"
	Path   string
	Level  string // Supports: "debug", "info", "warn", "error", "fatal"
	Daily  bool   // Supported driver: "file", Whether to generate a new log file every day
	Days   int    // Supported driver: "file", The number of days to keep the log file
}

var Log = &LogConfig{
	Driver: config.GetString("log.driver", "file"),
	Path:   config.GetString("log.path", "logs/"+str.ToSnakeCase(App.Name)+".log"),
	Level:  config.GetString("log.level", "info"),
	Daily:  config.GetBool("log.daily", true),
	Days:   config.GetInt("log.days", 15),
}
