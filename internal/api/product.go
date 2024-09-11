package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/gojwt/domain"
	"github.com/khairulharu/gojwt/dto"
	"github.com/khairulharu/gojwt/internal/util"
)

type apiProduct struct {
	productService domain.ProductService
}

func NewProduct(app *fiber.App, productService domain.ProductService, authMid fiber.Handler) {
	h := &apiProduct{
		productService: productService,
	}

	app.Get("product/", h.GetProduct)
	app.Get("products/", h.GetAllProduct)
	app.Post("product/", h.UpdateProduct)
	app.Post("product/insert", h.InsertProduct)
	app.Delete("delete/", h.DeleteProduct)
	app.Post("product/register", h.RegisterProduct)
	app.Get("product/validate", authMid, h.ValidateToken)
}

func (a apiProduct) GetProduct(ctx *fiber.Ctx) error {
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)
	res := a.productService.Find(ctx.Context(), int64(id))
	return ctx.Status(200).JSON(res)
}

func (a apiProduct) GetAllProduct(ctx *fiber.Ctx) error {
	res := a.productService.Get(ctx.Context())
	return ctx.Status(200).JSON(res)
}

func (a apiProduct) UpdateProduct(ctx *fiber.Ctx) error {
	var productRequ dto.ProductReq
	if err := ctx.BodyParser(&productRequ); err != nil {
		ctx.SendStatus(400)
	}
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)

	productRequ.ID = int64(id)

	res := a.productService.Update(ctx.Context(), productRequ)

	return ctx.Status(200).JSON(res)
}

func (a apiProduct) InsertProduct(ctx *fiber.Ctx) error {
	var product dto.ProductReq

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.SendStatus(400)
	}

	res := a.productService.Insert(ctx.Context(), product)

	return ctx.Status(200).JSON(res)
}

func (a apiProduct) DeleteProduct(ctx *fiber.Ctx) error {
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)

	res := a.productService.Delete(ctx.Context(), int64(id))

	return ctx.Status(200).JSON(res)
}

func (a apiProduct) RegisterProduct(ctx *fiber.Ctx) error {
	var product dto.ProductReq
	err := ctx.BodyParser(&product)
	if err != nil {
		ctx.Status(400).JSON(err.Error())
	}

	res := a.productService.Register(ctx.Context(), product)
	return ctx.Status(200).JSON(res)
}

func (a apiProduct) ValidateToken(ctx *fiber.Ctx) error {
	res := ctx.Locals("product").(*util.MyCustomClaims)
	product := dto.ProductRes{
		ID:   res.ID,
		Name: res.Name,
	}
	return ctx.Status(200).JSON(product)
}
