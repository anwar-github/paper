package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"papper/exception"
	"papper/middleware"
	"papper/service"
	"papper/transformer"
)

type Controller struct {
	Service     service.Interface
	Transformer transformer.Interface
}

type Interface interface {
	Disbursement(w http.ResponseWriter, r *http.Request)
}

func NewController(service service.Interface, transformer transformer.Interface) Interface {
	return Controller{
		Service:     service,
		Transformer: transformer,
	}
}

func (controller Controller) Disbursement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var request service.DisbursementRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// TODO: handle request validations
		// Middleware
		if err := middleware.Validation(request); err != nil {
			exceptions := exception.MapResponse(errors.New(exception.RequestValidation))

			w.WriteHeader(exceptions.GetHttpCode())
			json.NewEncoder(w).Encode(controller.Transformer.DisbursementRequestValidationTransform(err, exceptions))
			return
		}

		data, err := controller.Service.Disbursement(&request)
		exceptions := exception.MapResponse(err)

		w.WriteHeader(exceptions.GetHttpCode())
		json.NewEncoder(w).Encode(controller.Transformer.DisbursementTransform(data, exceptions))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w)
	}
}
