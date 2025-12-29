package handlers

import (
	"context"

	"github.com/ansufw/gomono/internal/apps/web/data"
	"github.com/ansufw/gomono/internal/config"
)

type Handler struct {
	ctxTempl  context.Context
	cfg       *config.Config
	templData *data.TemplGo
}

func New(ctx context.Context, cfg *config.Config, td *data.TemplGo) *Handler {

	// ctx = context.WithValue(ctx, "api-url", fmt.Sprintf("%s:%s", td.APIHost, td.APIPort))

	return &Handler{
		ctxTempl:  ctx,
		cfg:       cfg,
		templData: td,
	}
}
