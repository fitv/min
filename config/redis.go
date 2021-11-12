package config

import "github.com/fitv/min/core/config"

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Database int
}

var Redis = &RedisConfig{
	Host:     config.GetString("redis.host", "127.0.0.1"),
	Port:     config.GetInt("redis.port", 6379),
	Password: config.GetString("redis.password", ""),
	Database: 0,
}
