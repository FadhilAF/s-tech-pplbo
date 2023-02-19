package utils

import (
  "github.com/FadhilAF/s-tech-pplbo/internal/model"
)

func ToWebServiceResponse(message string, status int, data interface{}) model.WebServiceResponse {
	return model.WebServiceResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}
