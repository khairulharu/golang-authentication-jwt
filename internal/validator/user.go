package validator

import (
	"errors"

	"github.com/khairulharu/gojwt/dto"
)

func ValidateUserRequest(request dto.UserRequest) (validRequest dto.UserRequest, err error) {
	switch {
	case request.Name == (""):
		return dto.UserRequest{}, errors.New("validation Request: name must be define")
	case request.Username == (""):
		return dto.UserRequest{}, errors.New("validation Request: username are nil")
	case request.Password == (""):
		return dto.UserRequest{}, errors.New("validation Request: password mus define")
	default:
		return request, nil
	}
}
