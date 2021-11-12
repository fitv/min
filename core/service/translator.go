package service

import (
	"github.com/fitv/min/core/app"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Translator struct {
	Service
}

func (Translator) Register(app *app.Application) {
	en := en.New()
	zh := zh.New()
	uni := ut.New(en, zh, en)
	trans, _ := uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)

	for _, validation := range app.Validations {
		validation(validate, trans)
	}
	zh_translations.RegisterDefaultTranslations(validate, trans)
	app.Translator = trans
}
