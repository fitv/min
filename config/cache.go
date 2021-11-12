package config

import "github.com/fitv/min/util/str"

type CacheConfig struct {
	Driver   string
	Prefix   string
	Database int
}

var Cache = &CacheConfig{
	Driver:   "redis",
	Prefix:   str.ToSnakeCase(App.Name) + ":cache:",
	Database: 1,
}
