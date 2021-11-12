package lang

import (
	"fmt"

	"github.com/spf13/viper"
)

var v *viper.Viper

func init() {
	v = viper.New()
}

func Set(key string, value interface{}) {
	v.Set(key, value)
}

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
