package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func App() (*fiber.App, error) {
	app := fiber.New()

	Middleware(app)
	Route(app)

	return app, nil
}

func Run(apiApp *fiber.App) {
	if err := apiApp.Listen(":4444"); err != nil {
		log.Fatalf("error listen to port 4444: %v", err)
	}
}
