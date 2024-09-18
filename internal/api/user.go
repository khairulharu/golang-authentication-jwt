package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/gojwt/domain"
	"github.com/khairulharu/gojwt/dto"
)

type apiUser struct {
	userService domain.UserService
}

func NewUserApi(app *fiber.App, userService domain.UserService) {
	userHandler := &apiUser{
		userService: userService,
	}

	app.Post("/api/signup", userHandler.SignUp)
}

func (api *apiUser) SignUp(ctx *fiber.Ctx) error {
	var request dto.SignUpRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(dto.Response{
			Code:    400,
			Message: "Body Parser Error: might body dosn fill",
			Error:   err.Error(),
		})
	}

	response := api.userService.SignUp(ctx.Context(), request)

	return ctx.Status(int(response.Code)).JSON(response)
}
