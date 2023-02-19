package utils

import (
	"encoding/json"
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/common/validation"
	"github.com/FadhilAF/s-tech-pplbo/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func BindWith(ctx *gin.Context, i interface{}, binding binding.Binding) bool {
	err := ctx.ShouldBindWith(i, binding)

	if err == nil {
		return true
	}

	validate(ctx, err)
	return false
}

func BindJSONAndValidate(ctx *gin.Context, i interface{}) bool {
	err := ctx.ShouldBindJSON(i)

	if err == nil {
		return true
	}
	validate(ctx, err)
	return false
}

func BindURIAndValidate(ctx *gin.Context, i interface{}) bool {

	err := ctx.ShouldBindUri(i)
	if err == nil {
		return true
	}
	validate(ctx, err)
	return false
}

func validate(ctx *gin.Context, err error) {
	if errValidations, ok := err.(validator.ValidationErrors); ok {
		res := validation.HandleValidationErrors(errValidations)
		ctx.JSON(res.Status, res)
	} else if _, ok := err.(*json.UnmarshalTypeError); ok {
		ctx.JSON(http.StatusBadRequest, model.WebServiceResponse{
			Message: "Schema request tidak valid",
			Status:  http.StatusBadRequest,
			Data:    nil,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, ToWebServiceResponse("Internal server error", http.StatusInternalServerError, nil))
	}
}
