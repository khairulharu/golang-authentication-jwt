package test

import (
	"context"
	"reflect"
	"testing"

	"github.com/khairulharu/gojwt/domain"
	"github.com/khairulharu/gojwt/dto"
	"github.com/khairulharu/gojwt/internal/repository"
	"github.com/khairulharu/gojwt/internal/service"
)

func TestUserRepository(t *testing.T) {
	userTest := domain.User{
		Username: "test2",
		Password: "test1",
		Name:     "testgg",
	}

	userRepositoryTest := repository.NewUserRepository(dbGorm)
	t.Run("should be able to create new User", func(t *testing.T) {
		user, err := userRepositoryTest.Insert(context.Background(), &userTest)

		if user == (domain.User{}) {
			t.Error("error input new user")
		}

		if err != nil {
			t.Errorf("error input user: %v", err.Error())
		}
	})

	t.Run("should be able to find user using username", func(t *testing.T) {
		user, err := userRepositoryTest.FindByUsername(context.Background(), userTest.Username)

		if user == (domain.User{}) {
			t.Error("error input new user")
		}

		if err != nil {
			t.Errorf("error input user: %v", err.Error())
		}

		if !reflect.DeepEqual(user.Username, userTest.Username) {
			t.Error("error equaling username")
		}

		if !reflect.DeepEqual(user.Password, userTest.Password) {
			t.Error("error equaling username")
		}

		if !reflect.DeepEqual(user.Name, userTest.Name) {
			t.Error("error equaling username")
		}
	})

	t.Run("should be able to find user using username", func(t *testing.T) {
		err := userRepositoryTest.Delete(context.Background(), userTest.Username)

		if err != nil {
			t.Errorf("error input user: %v", err.Error())
		}
	})
}

func TestUserService(t *testing.T) {
	userRepositoryTest := repository.NewUserRepository(dbGorm)

	userServiceTest := service.NewUserService(userRepositoryTest)

	t.Run("should be to create new user with no invalid data", func(t *testing.T) {
		request := dto.SignUpRequest{
			Name:     "test",
			Username: "test",
			Password: "test",
		}

		response := userServiceTest.SignUp(context.Background(), request)

		if response.Code != 200 {
			t.Error(response)
		}
	})

	t.Run("should reject create data if have same username", func(t *testing.T) {
		request := dto.SignUpRequest{
			Name:     "ddd",
			Username: "ddd",
			Password: "ddd",
		}

		response := userServiceTest.SignUp(context.Background(), request)

		if response.Code != 400 {
			t.Error(response)
		}
	})
}
