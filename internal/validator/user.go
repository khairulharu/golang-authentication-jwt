package validator

import (
	"github.com/khairulharu/gojwt/dto"
)

func ValidateSignUpRequest(request dto.SignUpRequest) (validRequest dto.SignUpRequest, err error) {
	type validataeUserRequest struct {
		Username string `validate:"required,min=1,max=255"`
		Password string `validate:"required,min=1,max=255"`
		Name     string `validate:"required,min=1,max=255"`
	}

	var requestValidation = validataeUserRequest{
		Username: request.Username,
		Password: request.Password,
		Name:     request.Name,
	}

	if err := validate.Struct(requestValidation); err != nil {
		return dto.SignUpRequest{}, err
	}

	return dto.SignUpRequest(requestValidation), nil
}

func ValidateLogiInRequest(request dto.LogInRequest) (validRequest dto.LogInRequest, err error) {
	type validataeUserRequest struct {
		Username string `validate:"required,min=1,max=255"`
		Password string `validate:"required,min=1,max=255"`
	}

	var requestValidation = validataeUserRequest{
		Username: request.Username,
		Password: request.Password,
	}

	if err := validate.Struct(requestValidation); err != nil {
		return dto.LogInRequest{}, err
	}

	return dto.LogInRequest(requestValidation), nil
}
