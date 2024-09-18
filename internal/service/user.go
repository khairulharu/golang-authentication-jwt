package service

import (
	"context"

	"github.com/khairulharu/gojwt/domain"
	"github.com/khairulharu/gojwt/dto"
	"github.com/khairulharu/gojwt/internal/util"
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

func (u *userService) LogIn(ctx context.Context, request dto.LogInRequest) dto.Response {
	loginRequest, err := validator.ValidateLogiInRequest(request)
	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "Validate Error, Username or Password is required",
			Error:   err.Error(),
		}
	}

	isUserExist, err := u.userRepository.FindByUsername(ctx, loginRequest.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return dto.Response{
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}
	}

	if isUserExist == (domain.User{}) {
		return dto.Response{
			Code:    401,
			Message: "Error Username or Pasword Is Invalid",
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(isUserExist.Password), []byte(loginRequest.Password)); err != nil {
		return dto.Response{
			Code:    401,
			Message: "Error Username or Pasword Is Invalid",
			Error:   err.Error(),
		}
	}

	token, err := util.CreateToken(&isUserExist)
	if err != nil {
		return dto.Response{
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "Success",
		Data:    token,
	}
}

func (u *userService) SignUp(ctx context.Context, request dto.SignUpRequest) dto.Response {
	signUpRequest, err := validator.ValidateSignUpRequest(request)

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

func (u *userService) LogOut(ctx context.Context) dto.Response {
	panic("unimplemented")
}
