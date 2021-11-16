package lang

var (
	DefaultLocale string
	emptyTrans    = NewTranslator()
	transMap      = make(map[string]*Translator)
)

// Set sets the locale translator instance
func Set(locale string, translator *Translator) {
	transMap[locale] = translator
}

// Locale returns the translator instance by the given locale
func Locale(locale string) *Translator {
	trans, ok := transMap[locale]
	if ok {
		return trans
	}
	return emptyTrans
}

// Trans returns language translation by the given key
func Trans(key string, args ...interface{}) string {
	trans, ok := transMap[DefaultLocale]
	if ok {
		return trans.Trans(key, args...)
	}
	return emptyTrans.Trans(key, args...)
}
