package web

import (
	"context"

	"github.com/ansufw/gomono/internal/apps/web/data"
	"github.com/ansufw/gomono/internal/apps/web/handlers"
	"github.com/ansufw/gomono/internal/config"
	"github.com/gofiber/fiber/v2"
)

func App(cfg *config.Config) (*fiber.App, error) {
	ctx := context.Background()
	app := fiber.New()
	templData := data.New(cfg)

	handler := handlers.New(ctx, cfg, templData)

	Middleware(app)
	app.Use(handler.CtxTempl(templData))

	Route(app, handler)
	return app, nil
}

func Run(webApp *fiber.App) {
	// This function is no longer needed as we run from main app
}
