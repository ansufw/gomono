package handlers

import (
	"github.com/ansufw/gomono/views/pages"
	"github.com/gofiber/fiber/v2"
)

func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {

		c.Set("Content-Type", "text/html")
		return pages.Login().Render(c.Context(), c.Response().BodyWriter())
	}
}
