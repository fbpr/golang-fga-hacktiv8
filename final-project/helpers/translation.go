package helpers

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	v                                = validator.New()
	eng                              = en.New()
	uni      *ut.UniversalTranslator = ut.New(eng, eng)
	trans, _                         = uni.GetTranslator("en")
	_                                = en_translations.RegisterDefaultTranslations(v, trans)
)

func ValidateStruct(u interface{}) (err error) {
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
