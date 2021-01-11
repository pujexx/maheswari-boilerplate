package lib

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
)



func ValidateStruct(data interface{}) (bool, []ValidateError) {
	validate := validator.New()
	errs := []ValidateError{}
	if err := validate.Struct(data); err != nil {
		validatorErros := err.(validator.ValidationErrors)
		logrus.Println("validation", validatorErros)
		en := en.New()
		uni := ut.New(en, en)
		trans, _ := uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(validate, trans)
		for _, e := range validatorErros {
			errs = append(errs, ValidateError{
				Field: e.Field(),
				Error: e.Translate(trans),
			})
		}
		return false, errs
	}
	return true, errs
}
