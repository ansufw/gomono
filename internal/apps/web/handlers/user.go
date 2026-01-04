package handlers

import (
	"github.com/ansufw/gomono/views/templ/pages"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Login(c *fiber.Ctx) error {

	c.Set("Content-Type", "text/html")
	// templData := h.ctxTempl.Value("template-data").(*data.TemplGo)
	return pages.Login().Render(c.Context(), c.Response().BodyWriter())
}
