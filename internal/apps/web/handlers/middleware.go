package handlers

import (
	"github.com/ansufw/gomono/internal/apps/web/data"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CtxTempl(td *data.TemplGo) fiber.Handler {

	return func(c *fiber.Ctx) (err error) {

		c.Context().SetUserValue("template-data", td)

		return c.Next()
	}
}
