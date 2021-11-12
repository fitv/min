package config

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var v *viper.Viper

func init() {
	v = viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config file error: %w", err))
	}
}

// Set sets the value for the key in the config.
func Set(key string, value interface{}) {
	v.Set(key, value)
}

// Get gets the value for the key in the config.
func Get(key string, defaultValue ...interface{}) interface{} {
	if v.IsSet(key) {
		return v.Get(key)
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}

// GetStringSlice gets the string value for the key in the config.
func GetString(key string, defaultValue ...interface{}) string {
	return cast.ToString(Get(key, defaultValue...))
}

// GetInt gets the int value for the key in the config.
func GetInt(key string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(key, defaultValue...))
}

// GetBool gets the bool value for the key in the config.
func GetBool(key string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(key, defaultValue...))
}
