package middleware

import (
	"github.com/ansufw/gomono/internal/apps/web/data"
	"github.com/gofiber/fiber/v2"
)

func New(td *data.TemplGo) fiber.Handler {

	return func(c *fiber.Ctx) (err error) {

		c.Context().SetUserValue("template-data", "xxxxx")

		return c.Next()
	}
}
