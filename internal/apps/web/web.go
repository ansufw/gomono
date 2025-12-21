package web

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

func Run(webApp *fiber.App) {
	if err := webApp.Listen(":4444"); err != nil {
		log.Fatalf("error listen to port 4444: %v", err)
	}
}
