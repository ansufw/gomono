package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Middleware(app *fiber.App) {

	app.Use(recover.New())
	app.Use(logger.New())

}
