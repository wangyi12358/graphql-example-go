package validate

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
	Trans    ut.Translator
)

func JoinValidateErrors(err error) error {
	errs := err.(validator.ValidationErrors)
	var str string

	errsMap := errs.Translate(Trans)
	var stringSlice []string
	for _, v := range errsMap {
		stringSlice = append(stringSlice, v)
	}
	str = strings.Join(stringSlice, "，")
	return errors.New(str)
}

func Setup() {
	zh := zh.New()
	uni = ut.New(zh, zh)
	Trans, _ = uni.GetTranslator("zh")
	// TagNameFunc
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("name")
		if name != "" {
			return name
		}
		return field.Name
	})
	zh_translations.RegisterDefaultTranslations(Validate, Trans)
}
