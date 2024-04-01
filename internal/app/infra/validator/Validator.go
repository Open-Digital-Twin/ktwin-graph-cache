package validator

import (
	v10 "github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateStruct(interface{}) error
}

func NewValidator() Validator {
	return &validator{
		validator: *v10.New(),
	}
}

type validator struct {
	validator v10.Validate
}

func (v *validator) ValidateStruct(obj interface{}) error {
	return v.validator.Struct(obj)
}
