package api

import (
	"context"

	"github.com/ansufw/gomono/internal/config"
	"github.com/ansufw/gomono/internal/infrastructure/database/pg"
	"github.com/gofiber/fiber/v2"
)

func App(ctx context.Context, cfg *config.Config, db *pg.PG) (*fiber.App, error) {
	app := fiber.New()

	h := NewHandler(db)

	Middleware(app)
	Route(app, h)

	return app, nil
}

func Run(apiApp *fiber.App) {
	// This function is no longer needed as we run from main app
}
