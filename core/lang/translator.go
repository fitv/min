package lang

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewTranslator creates a new Translator
func NewTranslator() *Translator {
	return &Translator{viper: viper.New()}
}

// Translator is a viper wrapper
type Translator struct {
	viper *viper.Viper
}

// Set sets a translation
func (t *Translator) Set(key string, value interface{}) {
	t.viper.Set(key, value)
}

// Trans returns language translation by the given key
func (t *Translator) Trans(key string, args ...interface{}) string {
	if !t.viper.IsSet(key) {
		return key
	}

	value := t.viper.GetString(key)

	if len(args) > 0 {
		return fmt.Sprintf(value, args...)
	}
	return value
}
