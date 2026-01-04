package web

import (
	"github.com/ansufw/gomono/internal/shared/enum"
	"github.com/ansufw/gomono/internal/shared/zl"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Middleware(app *fiber.App) {
	app.Use(recover.New())

	logger := zl.Logger(enum.ModeDev, enum.ServiceWeb)
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))
}
