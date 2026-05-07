package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {

	validationMap := map[string]string{}

	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {

		for _, fieldError := range validationErrors {

			field := strings.ToLower(fieldError.Field())

			switch fieldError.Tag() {

			case "required":
				validationMap[field] = field + " is required"

			case "email":
				validationMap[field] = "invalid email format"

			case "min":
				validationMap[field] = field + " is too short"

			case "max":
				validationMap[field] = field + " is too long"

			default:
				validationMap[field] = "invalid value"
			}
		}
	}

	return validationMap
}
