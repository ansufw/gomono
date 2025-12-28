package api

import (
	_ "github.com/ansufw/gomono/docs/api" // Import generated docs to register Swagger
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Route(app *fiber.App, h *Handler) {
	api := app.Group("/api")

	api.Get("/swagger/*", swagger.HandlerDefault)

	// Add root API handler for /api endpoint
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API is running",
			"version": "1.0.0",
		})
	})

	api.Get("/users", h.ListUsers)

}
