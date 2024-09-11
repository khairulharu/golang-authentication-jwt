package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/khairulharu/gojwt/internal/api"
	"github.com/khairulharu/gojwt/internal/components"
	"github.com/khairulharu/gojwt/internal/config"
	"github.com/khairulharu/gojwt/internal/repository"
	"github.com/khairulharu/gojwt/internal/service"
)

func main() {
	cnf := config.New()
	dbConnection := components.NewDatabaseConnection(cnf)

	userRepository := repository.NewUserRepository(dbConnection)

	userServie := service.NewUserService(userRepository)

	// authMid := middelware.Authenticate()

	app := fiber.New()
	app.Use(logger.New())

	api.NewUserApi(app, userServie)

	app.Listen(cnf.SRV.Host + ":" + cnf.SRV.Port)
}
