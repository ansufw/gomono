package web

import (
	"fmt"

	"github.com/ansufw/gomono/internal/apps/web/handlers"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, h *handlers.Handler) {
	app.Static("/static", "./public")

	// Add favicon route for debugging
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		fmt.Println("DEBUG: Favicon.ico route accessed")
		return c.Status(404).SendString("Favicon not found")
	})

	app.Get("/", h.Home)
	app.Get("/login", h.Login)
	// app.Post("/login", handlers.Login())
}
