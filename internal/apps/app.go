package apps

import (
	"context"
	"fmt"
	"log"

	"github.com/ansufw/gomono/internal/apps/api"
	"github.com/ansufw/gomono/internal/apps/web"
	"github.com/ansufw/gomono/internal/config"
	"github.com/ansufw/gomono/internal/infrastructure/database/pg"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Fiber  *fiber.App
	config *config.Config
}

func Wire() (*App, error) {

	ctx := context.Background()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	database := pg.New(ctx, cfg.DB.Url)

	// Create Fiber apps
	apiApp, err := api.App(ctx, cfg, database)
	if err != nil {
		return nil, err
	}

	webApp, err := web.App(cfg)
	if err != nil {
		return nil, err
	}

	// Mount web app to api app
	apiApp.Mount("/", webApp)

	return &App{
		Fiber:  apiApp,
		config: cfg,
	}, nil
}

func Run(app *App) {
	// defer app.DB.Close()

	// DEBUG: Log configuration values
	log.Printf("API Config - Host: '%s', Port: %d", app.config.Api.Host, app.config.Api.Port)

	// BUG FIX: Use Port instead of Host
	addr := fmt.Sprintf(":%d", app.config.Api.Port)
	log.Printf("Starting server on %s", addr)

	if err := app.Fiber.Listen(addr); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
