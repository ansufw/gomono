package handlers

import (
	"github.com/ansufw/gomono/views/pages"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Home(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	// c.Context().SetUserValue("template-data", h.templData)
	return pages.Home().Render(c.Context(), c.Response().BodyWriter())
}
