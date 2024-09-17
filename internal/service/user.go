package service

import (
	"context"

	"github.com/khairulharu/gojwt/domain"
	"github.com/khairulharu/gojwt/dto"
	"github.com/khairulharu/gojwt/internal/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	checkUserIsSameUsername, err := u.userRepository.FindByUsername(ctx, signUpRequest.Username)

	if err != nil && err != gorm.ErrRecordNotFound {
		return dto.Response{
			Code:    500,
			Message: "Error When Check new User have a same username",
			Error:   err.Error(),
		}
	}

	if checkUserIsSameUsername != (domain.User{}) {
		return dto.Response{
			Code:    400,
			Message: "Chechker Username: Username or Password invalid",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), 10)

	if err != nil {
		return dto.Response{
			Code:    500,
			Message: "Internal Server Error: While Generate Hashing Password",
			Error:   err.Error(),
		}
	}

	user := domain.User{
		Username: signUpRequest.Username,
		Password: string(hashedPassword),
		Name:     signUpRequest.Name,
	}

	user, err = u.userRepository.Insert(ctx, &user)

	if err != nil {
		return dto.Response{
			Code:    500,
			Message: "Error When Inserted Data",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "Success Create User",
		Data: dto.UserResponse{
			Username: user.Username,
			Name:     user.Name,
		},
	}
}
