package domain

import (
	"context"

	"github.com/khairulharu/gojwt/dto"
)

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Name     string `db:"name"`
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	FindByUsername(ctx context.Context, username string) (User, error)
	Delete(ctx context.Context, username string) error
}

type UserService interface {
	SignUp(ctx context.Context, request dto.UserRequest) dto.Response
	LogIn(ctx context.Context, request dto.UserRequest) dto.Response
	LogOut(ctx context.Context) dto.Response
}
