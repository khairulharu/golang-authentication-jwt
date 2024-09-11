package middelware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/gojwt/internal/util"
)

func Authenticate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if token == "" {
			return ctx.SendStatus(401)
		}

		product, err := util.ValidateToken(token)
		if err != nil {
			return ctx.Status(400).JSON(err.Error())
		}
		ctx.Locals("product", product)
		return ctx.Next()
	}
}
