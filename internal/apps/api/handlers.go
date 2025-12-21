package api

import "github.com/gofiber/fiber/v2"

func GetUser(c *fiber.Ctx) error {
	// Simple user list for testing without database connection
	users := []map[string]interface{}{
		{
			"id":         1,
			"first_name": "John",
			"last_name":  "Doe",
			"email":      "john@example.com",
		},
		{
			"id":         2,
			"first_name": "Jane",
			"last_name":  "Smith",
			"email":      "jane@example.com",
		},
	}

	return c.JSON(users)
}
