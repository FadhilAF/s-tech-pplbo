package validation

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func registerCustomErrors() {
	customErrors["required"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("%s tidak boleh kosong", translatedFieldName)
	}

	customErrors["startswith"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("%s harus berawalan %s", translatedFieldName, field.Value())
	}

	customErrors["uuid"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("%s harus dengan format UUID", translatedFieldName)
	}

	customErrors["max"] = func(field validator.FieldError, translatedFieldName string) string {
		fieldType := field.Kind()
		if fieldType == reflect.Int {
			return fmt.Sprintf("%s tidak boleh lebih dari %s", translatedFieldName, field.Param())
		}
		return fmt.Sprintf("%s tidak boleh melebihi %s karakter", translatedFieldName, field.Param())
	}

	customErrors["numeric"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("%s harus numeric", translatedFieldName)
	}
}
