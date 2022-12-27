package helpers

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func TranslateError(v *validator.Validate, u interface{}) (err error) {
	var uni *ut.UniversalTranslator

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.Replace(strings.SplitN(fld.Tag.Get("json"), ",", 2)[0], "_", " ", -1)
		if name == "-" {
			return ""
		}

		return name
	})

	v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "Your {0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	var errStrings []string
	e := v.Struct(u)
	if e != nil {
		for _, e := range e.(validator.ValidationErrors) {
			errStrings = append(errStrings, e.Translate(trans))
		}

		return errors.New(strings.Join(errStrings[:], ";"))
	}

	return
}
