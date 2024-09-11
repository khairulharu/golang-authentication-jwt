package domain

import (
	"context"

	"github.com/khairulharu/gojwt/dto"
)

type Product struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type ProductRepository interface {
	GetAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id int64) (Product, error)
	FindByName(ctx context.Context, name string) (Product, error)
	UpdateByID(ctx context.Context, product *Product) error
	InsertProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, product Product) error
}

type ProductService interface {
	Get(ctx context.Context) dto.Response
	Find(ctx context.Context, id int64) dto.Response
	Update(ctx context.Context, req dto.ProductReq) dto.Response
	Insert(ctx context.Context, req dto.ProductReq) dto.Response
	Delete(ctx context.Context, id int64) dto.Response
	Register(ctx context.Context, req dto.ProductReq) dto.Response
}
