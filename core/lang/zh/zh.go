package zh

import (
	"github.com/spf13/viper"
)

var lang *viper.Viper

func init() {
	lang = viper.New()
}

func Set(key string, value interface{}) {
	lang.Set(key, value)
}

func Lang() *viper.Viper {
	return lang
}
