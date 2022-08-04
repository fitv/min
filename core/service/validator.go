package service

import (
	"regexp"

	"github.com/fitv/min/core/app"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	regexpMobile = regexp.MustCompile("^1[3-9][0-9]{9}$")
)

type Validator struct {
	Service
}

func (Validator) Register(app *app.Application) {
	app.AddValidation(func(validate *validator.Validate, trans ut.Translator) {
		validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
			return regexpMobile.MatchString(fl.Field().Interface().(string))
		})

		validate.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
			return ut.Add("mobile", app.Lang.Trans("validation.mobile"), false)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	})
}
