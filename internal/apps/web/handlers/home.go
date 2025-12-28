package handlers

import (
	"github.com/ansufw/gomono/views/pages"
	"github.com/gofiber/fiber/v2"
)

func Home() fiber.Handler {
	return func(c *fiber.Ctx) error {

		c.Set("Content-Type", "text/html")
		return pages.Home().Render(c.Context(), c.Response().BodyWriter())
	}
}
