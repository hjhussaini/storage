package models

import (
	"strings"

	"github.com/go-playground/validator"
)

type Validation struct {
	validate *validator.Validate
}

func (validation *Validation) Validate(object interface{}) FieldErrors {
	err := validation.validate.Struct(object)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)
	var fields FieldErrors

	for _, err := range errs {
		field := FieldError{err.(validator.FieldError)}
		fields = append(fields, field)
	}

	return fields
}

func validatePath(field validator.FieldLevel) bool {
	path := field.Field().String()

	return strings.HasPrefix(path, "/") && !strings.HasSuffix(path, "/")
}

// NewValidation creates a new validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("path", validatePath)

	return &Validation{validate}
}
