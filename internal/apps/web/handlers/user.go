package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Redirect("/")
	}
}
