package transformer

import (
	"papper/exception"
	"papper/middleware"
	"papper/service"
)

type Transformer struct {
}

type Interface interface {
	DisbursementTransform(response *service.DisbursementResponse, exceptionInterface exception.MessageExceptionInterface) *Success
	DisbursementRequestValidationTransform(middleware []*middleware.ErrorValidationResponse, exception exception.MessageExceptionInterface) *RequestValidation
}

func NewTransformer() Interface {
	return Transformer{}
}

type Success struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RequestValidation struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (t Transformer) Json(data interface{}, exception exception.MessageExceptionInterface) *Success {
	return &Success{
		Code: exception.GetResponseCode(),
		// TODO: handle error message with json list
		Message: exception.GetResponseCode(),
		Data:    data,
	}
}

func (t Transformer) HandleRequestValidation(validation []*middleware.ErrorValidationResponse, exception exception.MessageExceptionInterface) *RequestValidation {
	return &RequestValidation{
		Code:    exception.GetResponseCode(),
		Message: exception.GetResponseCode(),
		Data:    validation,
	}
}
