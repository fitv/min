package config

import (
	"github.com/fitv/min/core/config"
)

type AppConfig struct {
	Name   string
	Addr   string
	Port   int
	Debug  bool
	Locale string
}

var App = &AppConfig{
	Name:   config.GetString("app.name", "MIN"),
	Addr:   config.GetString("app.addr", "127.0.0.1"),
	Port:   config.GetInt("app.port", 3000),
	Debug:  config.GetBool("app.debug", false),
	Locale: config.GetString("app.locale", "zh"),
}
