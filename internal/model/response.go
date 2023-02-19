package model

type WebServiceResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type ValidationErrorWebServiceResponse struct {
	WebServiceResponse
	DetailErrors []string `json:"errors"`
}
