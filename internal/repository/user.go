package repository

import (
	"context"

	"github.com/khairulharu/gojwt/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	dbGorm *gorm.DB
}

func NewUserRepository(dbGorm *gorm.DB) domain.UserRepository {
	return &userRepository{
		dbGorm: dbGorm,
	}
}

// Insert implements domain.UserRepository.
func (u *userRepository) Insert(ctx context.Context, user *domain.User) (domain.User, error) {
	err := u.dbGorm.Debug().WithContext(ctx).Table("users").Create(&user).Error
	return *user, err
}

// Delete implements domain.UserRepository.
func (u *userRepository) Delete(ctx context.Context, username string) error {
	err := u.dbGorm.Debug().WithContext(ctx).Table("users").Delete(&domain.User{Username: username}).Error
	return err
}

// FindByUsername implements domain.UserRepository.
// Change This
func (u *userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	err = u.dbGorm.Debug().WithContext(ctx).Table("users").Where("username = ?", username).First(&user).Error
	return
}

//Mendapatkan nil pointer dereferences, coba perbaiki!!
//Dan pada halaman findbYusername dan bagian service
