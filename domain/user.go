package domain

import (
	"context"

	"github.com/khairulharu/gojwt/dto"
)

type User struct {
	Username string `gorm:"primaryKey; size: 255; not null; unique"`
	Password string `gorm:"not null; size: 255"`
	Name     string `gorm:"size: 255"`
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Delete(ctx context.Context, username string) error
}

type UserService interface {
	SignUp(ctx context.Context, request dto.UserRequest) dto.Response
	LogIn(ctx context.Context, request dto.UserRequest) dto.Response
	LogOut(ctx context.Context) dto.Response
}
