package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationErrors struct {
	Errors []string `json:"errors"`
}

func New() *validator.Validate {
	v := validator.New()

	// register field names specified in the JSON struct tags instead of go struct fields
	v.RegisterTagNameFunc(func(sf reflect.StructField) string {
		n := strings.SplitN(sf.Tag.Get("json"), ",", 2)[0]
		if n == "-" {
			return ""
		}
		return n
	})

	return v
}

func ErrResponse(err error) *ValidationErrors {
	if valErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ValidationErrors{
			Errors: make([]string, len(valErrors)),
		}

		// loop through the validation errors
		for i, fieldErr := range valErrors {
			// get the field name and error message
			fieldName := fieldErr.StructField()
			errorMessage := fieldErr.ActualTag()

			switch fieldErr.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is required", fieldName)
			case "oneof":
				resp.Errors[i] = fmt.Sprintf("%s must be one of: %s", fieldName, fieldErr.Param())
			default:
				resp.Errors[i] = fmt.Sprintf("%s %s", fieldName, errorMessage)
			}
		}

		return &resp
	}

	return nil
}
