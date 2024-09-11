package service

import (
	"context"

	"github.com/khairulharu/gojwt/domain"
	"github.com/khairulharu/gojwt/dto"
	"github.com/khairulharu/gojwt/internal/validator"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// LogIn implements domain.UserService.
func (u *userService) LogIn(ctx context.Context, request dto.UserRequest) dto.Response {
	panic("unimplemented")
}

// LogOut implements domain.UserService.
func (u *userService) LogOut(ctx context.Context) dto.Response {
	panic("unimplemented")
}

// SignUp implements domain.UserService.
func (u *userService) SignUp(ctx context.Context, request dto.UserRequest) dto.Response {
	signUpRequest, err := validator.ValidateUserRequest(request)

	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "Validation Error",
			Error:   err.Error(),
		}
	}

	user := domain.User{
		Username: signUpRequest.Username,
		Password: signUpRequest.Password,
		Name:     signUpRequest.Name,
	}

	if err := u.userRepository.Insert(ctx, &user); err != nil {
		return dto.Response{
			Code:    500,
			Message: "Error When Inserted Data",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "Success User Created",
	}
}
