package requests

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.Split(field.Tag.Get("json"), ",")[0]
		if name == "" || name == "-" {
			return field.Name
		}
		return name
	})
}

func formatValidationErrors(err error) map[string]string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{
			"error": "invalid request",
		}
	}

	errors := make(map[string]string)

	for _, fieldErr := range validationErrors {
		switch fieldErr.Tag() {
		case "required":
			errors[fieldErr.Field()] = fieldErr.Field() + " is required"
		case "email":
			errors[fieldErr.Field()] = fieldErr.Field() + " must be a valid email"
		case "min":
			errors[fieldErr.Field()] = fieldErr.Field() + " is too short"
		case "max":
			errors[fieldErr.Field()] = fieldErr.Field() + " is too long"
		default:
			errors[fieldErr.Field()] = fieldErr.Field() + " is invalid"
		}
	}

	return errors
}
