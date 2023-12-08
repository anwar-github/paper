package transformer

import (
	"papper/exception"
	"papper/middleware"
	"papper/service"
	"reflect"
	"time"
)

type DisbursementTransformer struct {
	OrderID   string    `json:"order_id"`
	RequestID string    `json:"request_id"`
	Success   bool      `json:"success"`
	Time      time.Time `json:"time"`
}

func (t Transformer) DisbursementTransform(response *service.DisbursementResponse, exception exception.MessageExceptionInterface) *Success {
	if reflect.ValueOf(response.Disbursement).IsNil() {
		return t.Json(response.Request, exception)
	}

	return t.Json(&DisbursementTransformer{
		OrderID:   response.Disbursement.Order,
		RequestID: response.Disbursement.Request,
		Success:   response.Disbursement.Success,
		Time:      response.Disbursement.CreatedAt,
	}, exception)
}

func (t Transformer) DisbursementRequestValidationTransform(middleware []*middleware.ErrorValidationResponse, exception exception.MessageExceptionInterface) *RequestValidation {
	return t.HandleRequestValidation(middleware, exception)
}
