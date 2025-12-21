package apps

import (
	"log"

	"github.com/ansufw/gomono/internal/apps/api"
	"github.com/ansufw/gomono/internal/apps/web"
	"github.com/gofiber/fiber/v2"
)

func Wire() *fiber.App {
	apiApp, _ := api.App()
	webApp, _ := web.App()

	apiApp.Mount("/", webApp)

	return apiApp
}

func Run(app *fiber.App) {
	if err := app.Listen(":4444"); err != nil {
		log.Fatalf("error listen to port 4444: %v", err)
	}
}
