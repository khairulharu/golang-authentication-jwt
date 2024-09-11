package repository

import (
	"context"

	"github.com/khairulharu/gojwt/domain"
	"gorm.io/gorm"
)

type repositoryProduct struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) domain.ProductRepository {
	return &repositoryProduct{
		db: db,
	}
}

func (r repositoryProduct) GetAll(ctx context.Context) (product []domain.Product, err error) {
	err = r.db.Debug().WithContext(ctx).Table("products").Order("id").Find(&product).Error
	return
}

func (r repositoryProduct) FindByID(ctx context.Context, id int64) (product domain.Product, err error) {
	err = r.db.Debug().WithContext(ctx).Table("products").Where("id=?", id).First(&product).Error
	return
}

func (r repositoryProduct) FindByName(ctx context.Context, name string) (product domain.Product, err error) {
	err = r.db.Debug().WithContext(ctx).Table("products").Where("name = ?", name).First(&product).Error
	return
}

func (r repositoryProduct) UpdateByID(ctx context.Context, product *domain.Product) error {
	err := r.db.Debug().WithContext(ctx).Table("products").Updates(&product).Error
	return err
}

func (r repositoryProduct) InsertProduct(ctx context.Context, product *domain.Product) error {
	err := r.db.Debug().WithContext(ctx).Table("products").Create(&product).Error
	return err
}

func (r repositoryProduct) DeleteProduct(ctx context.Context, product domain.Product) error {
	err := r.db.Debug().WithContext(ctx).Delete(&product).Error
	return err
}
