package validator

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validate *validator.Validate
}

func NewCustomValidator(validate *validator.Validate) *CustomValidator {
	return &CustomValidator{validate: validate}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validate.Struct(i); err != nil {
		return err
	}
	return nil
}
