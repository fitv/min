package config

import (
	"github.com/fitv/min/core/config"
)

type DatabaseConfig struct {
	Driver    string
	Host      string
	Port      int
	Database  string
	User      string
	Password  string
	Charset   string
	Collation string
	Timeout   string // Dial Timeout, See https://github.com/go-sql-driver/mysql#timeout
	Debug     bool
}

var Database = &DatabaseConfig{
	Driver:    config.GetString("database.driver", "mysql"),
	Host:      config.GetString("database.host", "127.0.0.1"),
	Port:      config.GetInt("database.port", 3306),
	Database:  config.GetString("database.database"),
	User:      config.GetString("database.user"),
	Password:  config.GetString("database.password"),
	Charset:   config.GetString("database.charset", "utf8mb4"),
	Collation: config.GetString("database.collation", "utf8mb4_unicode_ci"),
	Timeout:   config.GetString("database.timeout", "5s"),
	Debug:     config.GetBool("database.debug", false),
}
