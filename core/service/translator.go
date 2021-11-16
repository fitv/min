package service

import (
	"fmt"

	"github.com/fitv/min/config"
	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/lang"
	_ "github.com/fitv/min/lang/en"
	_ "github.com/fitv/min/lang/zh"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Translator struct {
	Service
}

func (Translator) Register(app *app.Application) {
	var trans ut.Translator
	validate := binding.Validator.Engine().(*validator.Validate)

	switch config.App.Locale {
	case "en":
		en := en.New()
		uni := ut.New(en, en)
		trans, _ = uni.GetTranslator("en")

		for _, validation := range app.Validations {
			validation(validate, trans)
		}
		en_translations.RegisterDefaultTranslations(validate, trans)
	case "zh":
		en := en.New()
		zh := zh.New()
		uni := ut.New(en, zh, en)
		trans, _ = uni.GetTranslator("zh")
		for _, validation := range app.Validations {
			validation(validate, trans)
		}
		zh_translations.RegisterDefaultTranslations(validate, trans)
	default:
		panic(fmt.Errorf("unsupported locale: %s", config.App.Locale))
	}

	app.Translator = trans
	lang.DefaultLocale = config.App.Locale
}
