package handlers

import (
	"github.com/ansufw/gomono/views/pages"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// Handle GET request - show login page
	if c.Method() == "GET" {
		c.Set("Content-Type", "text/html")
		return pages.Login().Render(c.Context(), c.Response().BodyWriter())
	}

	// Handle POST request - process login form
	email := c.FormValue("email")
	password := c.FormValue("password")

	// TODO: Implement actual authentication logic
	// For now, just check if email and password are provided
	if email == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	// TODO: Validate credentials against database
	// For now, redirect to home on successful login
	return c.Redirect("/")
}
