package validation

import (
	"fmt"
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/go-playground/validator/v10"
)

func debugFieldError(v validator.FieldError) {
	fmt.Println("ini field", v.Field())
	fmt.Println("ini kind", v.Kind())
	fmt.Println("ini actual tag", v.ActualTag())
	fmt.Println("ini tag", v.Tag())
	fmt.Println("ini value", v.Value())
	fmt.Println("ini type", v.Type())
	fmt.Println("ini namespace", v.Namespace())
	fmt.Println("ini param", v.Param())
	fmt.Println("ini struct field", v.StructField())
	fmt.Println("ini struct namespace", v.StructNamespace())
}

func HandleValidationErrors(errs validator.ValidationErrors) (res model.ValidationErrorWebServiceResponse) {
	res.Status = http.StatusUnprocessableEntity
	var finalErrors []string
	for _, v := range errs {
		// debugFieldError(v)
		translateFunction, ok := customErrors[v.Tag()]
		if ok {
			translatedFieldName, found := customMessages[v.Field()]
			if !found {
				translatedFieldName = v.Field()
			}
			finalErrors = append(
				finalErrors,
				translateFunction(v, translatedFieldName),
			)
		}
	}
	if len(finalErrors) > 0 {
		res.Message = finalErrors[0]
	}

	res.DetailErrors = finalErrors
	res.Data = nil
	return res
}
