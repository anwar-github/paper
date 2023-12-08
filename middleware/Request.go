package middleware

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type ErrorValidationResponse struct {
	Field   string `json:"field" bson:"field"`
	Message string `json:"message" bson:"message"`
	Value   string `json:"value" bson:"value"`
}

func Validation(req interface{}) []*ErrorValidationResponse {
	var validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	var errors []*ErrorValidationResponse
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValidationResponse
			element.Field = err.Field()
			element.Message = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
