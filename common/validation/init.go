package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var customErrors map[string]func(field validator.FieldError, translatedFieldName string) string

var customMessages map[string]string

func InitValidations(validate *validator.Validate) {
	customMessages = map[string]string{}
	registerCustomMessages()

	customErrors = map[string]func(validator.FieldError, string) string{}
	registerCustomErrors()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
