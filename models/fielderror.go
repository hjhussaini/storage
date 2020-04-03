package models

import (
	"fmt"

	"github.com/go-playground/validator"
)

// FieldError wraps the validators FieldError
type FieldError struct {
	validator.FieldError
}

func (field FieldError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Failed validation for '%s' on the '%s' tag",
		field.Namespace(),
		field.Field(),
		field.Tag(),
	)
}

// FieldErrors is a collection of FieldError
type FieldErrors []FieldError

// FieldErrors converts the slice into a string slice
func (fields FieldErrors) Errors() []string {
	var errors []string

	for _, field := range fields {
		errors = append(errors, field.Error())
	}

	return errors
}
