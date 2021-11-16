package en

import (
	"github.com/fitv/min/core/lang"
)

var trans *lang.Translator

func init() {
	trans = lang.NewTranslator()

	lang.Set("en", trans)
}

func Set(key string, value interface{}) {
	trans.Set(key, value)
}
