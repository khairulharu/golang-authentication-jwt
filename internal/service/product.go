package service

import (
	"context"

	"github.com/khairulharu/gojwt/domain"
	"github.com/khairulharu/gojwt/dto"
	"github.com/khairulharu/gojwt/internal/util"
)

type serviceProduct struct {
	productRepository domain.ProductRepository
}

func NewProduct(productRepository domain.ProductRepository) domain.ProductService {
	return &serviceProduct{
		productRepository: productRepository,
	}
}

func (s serviceProduct) Get(ctx context.Context) dto.Response {
	products, err := s.productRepository.GetAll(ctx)

	if err != nil {
		return dto.Response{
			Code:    401,
			Message: "ERROR",
			Error:   err.Error(),
		}
	}

	var productsRes []dto.ProductRes

	for _, v := range products {
		productsRes = append(productsRes, dto.ProductRes{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
		Data:    productsRes,
	}
}

func (s serviceProduct) Find(ctx context.Context, id int64) dto.Response {
	product, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "ERROR",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
		Data: dto.ProductRes{
			ID:   product.ID,
			Name: product.Name,
		},
	}
}

func (s serviceProduct) Update(ctx context.Context, req dto.ProductReq) dto.Response {
	_, err := s.productRepository.FindByID(ctx, req.ID)
	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "ERROR",
			Error:   err.Error(),
		}
	}
	var product domain.Product
	product.ID = req.ID
	product.Name = req.Name

	err = s.productRepository.UpdateByID(ctx, &product)

	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "ERROR",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
	}
}

func (s serviceProduct) Insert(ctx context.Context, req dto.ProductReq) dto.Response {
	var product domain.Product
	product.ID = req.ID
	product.Name = req.Name

	resProduct, _ := s.productRepository.FindByName(ctx, product.Name)
	if resProduct.Name == product.Name {
		return dto.Response{
			Code:    400,
			Message: "name error try another",
		}
	}

	err := s.productRepository.InsertProduct(ctx, &product)
	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "ERROR",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
	}
}

func (s serviceProduct) Delete(ctx context.Context, id int64) dto.Response {
	resProduct, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "errorrr",
			Error:   err.Error(),
		}
	}

	err = s.productRepository.DeleteProduct(ctx, resProduct)
	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "errorrr",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
	}
}

func (s serviceProduct) Register(ctx context.Context, req dto.ProductReq) dto.Response {
	product, err := s.productRepository.FindByName(ctx, req.Name)
	if err != nil {
		return dto.Response{
			Code:    401,
			Message: "name of product not found",
			Error:   err.Error(),
		}
	}

	token, err := util.CreateToken(&product)

	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "EROROR TOKEN RESPON",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
		Data:    token,
	}
}
