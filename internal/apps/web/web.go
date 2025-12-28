package web

import (
	"github.com/ansufw/gomono/internal/config"
	"github.com/gofiber/fiber/v2"
)

func App(cfg *config.Config) (*fiber.App, error) {
	app := fiber.New()

	Middleware(app)
	Route(app)
	return app, nil
}

func Run(webApp *fiber.App) {
	// This function is no longer needed as we run from main app
}
