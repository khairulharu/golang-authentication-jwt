package validator

import (
	"github.com/khairulharu/gojwt/dto"
)

func ValidateUserRequest(request dto.UserRequest) (validRequest dto.UserRequest, err error) {
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
		return dto.UserRequest{}, err
	}

	return dto.UserRequest(requestValidation), nil
}
