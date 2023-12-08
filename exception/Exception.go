package exception

import (
	"errors"
	"net/http"
)

const InsufficientBalance = "INSUFFICIENT_BALANCE"
const DuplicateRequest = "DUPLICATE"
const InternalError = "INTERNAL_SERVER_ERROR"
const Ok = "OK"
const RequestValidation = "REQUEST_VALIDATION"

var (
	mapExceptionList = map[string]MessageExceptionInterface{
		InsufficientBalance: NewMessage(
			InsufficientBalance,
			http.StatusUnprocessableEntity,
			false),
		DuplicateRequest: NewMessage(
			DuplicateRequest,
			http.StatusUnprocessableEntity,
			false),
		Ok: NewMessage(
			Ok,
			http.StatusOK,
			false),
		InternalError: NewMessage(
			InternalError,
			http.StatusInternalServerError,
			true),
		RequestValidation: NewMessage(
			RequestValidation,
			http.StatusUnprocessableEntity,
			false),
	}
)

type MessageException struct {
	Service    string
	HttpCode   int
	ErrorCode  string
	Reportable bool
}

func NewMessage(errorCode string, httpCode int, IsReportable bool) MessageExceptionInterface {
	return MessageException{
		ErrorCode:  errorCode,
		HttpCode:   httpCode,
		Reportable: IsReportable,
	}
}

type MessageExceptionInterface interface {
	GetHttpCode() int
	GetResponseCode() string
	IsReportable() bool
}

func (exception MessageException) GetHttpCode() int {
	return exception.HttpCode
}

func (exception MessageException) GetResponseCode() string {
	return exception.ErrorCode
}

func (exception MessageException) IsReportable() bool {
	return exception.Reportable
}

func MapResponse(err error) MessageExceptionInterface {
	if err == nil {
		err = errors.New(Ok)
	}
	if value, ok := mapExceptionList[err.Error()].(MessageExceptionInterface); ok {
		return value
	}
	return mapExceptionList[InternalError]
}

type Exception struct {
	Code string `json:"code"`
}
