package helpers

import (
	"github.com/cryptosalamander/gorm_crud_example/dtos"
	"github.com/cryptosalamander/gorm_crud_example/langs"
	"github.com/go-playground/validator"
)

func GenerateValidationResponse(err error) (response dtos.ValidationResponse) {
	response.Success = false

	var validations []dtos.Validation

	validationErrors := err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		field, rule := value.Field, value.Tag

		validation := dtos.Validation{Field: field, Message: langs.GenerateValidationMessage(field, rule)}
		validations = append(validations, validation)
	}

	response.Validations = validations

	return response
}
