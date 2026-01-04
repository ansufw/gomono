package api

import (
	"github.com/ansufw/gomono/internal/shared/enum"
	"github.com/ansufw/gomono/internal/shared/zl"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Middleware(app *fiber.App) {

	app.Use(recover.New())

	logger := zl.Logger(enum.ModeDev, enum.ServiceAPI)

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, hx-current-url, hx-request, hx-trigger, hx-target, hx-current-url",
	}))

}
