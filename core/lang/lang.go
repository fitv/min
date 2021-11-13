package lang

import (
	"fmt"

	"github.com/spf13/viper"
)

var v *viper.Viper

// Set define default language instance
func Set(lang *viper.Viper) {
	v = lang
}

// Trans returns language translation by the given key
func Trans(key string, args ...interface{}) string {
	if !v.IsSet(key) {
		return key
	}

	value := v.GetString(key)

	if len(args) > 0 {
		return fmt.Sprintf(value, args...)
	}
	return value
}
