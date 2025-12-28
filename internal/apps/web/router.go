package web

import (
	"github.com/ansufw/gomono/internal/apps/web/handlers"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Static("/static", "./public")

	app.Get("/", handlers.Home())
	app.Get("/login", handlers.Login())
	// app.Post("/login", handlers.Login())
}
