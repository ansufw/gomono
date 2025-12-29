package data

import (
	"context"
	"fmt"
	"log"

	"github.com/ansufw/gomono/internal/config"
)

// TemplateData holds common data that can be passed to all templates
type TemplGo struct {
	APIHost string
	APIPort int
	WebHost string
	WebPort int
	// Add any other global template data here
}

// NewTemplateData creates a new TemplateData instance from config
func New(cfg *config.Config) *TemplGo {
	return &TemplGo{
		APIHost: cfg.Api.Host,
		APIPort: cfg.Api.Port,
		WebHost: cfg.Web.Host,
		WebPort: cfg.Web.Port,
	}
}

func GetCtx[T any](ctx context.Context, key string) (T, bool) {
	var zero T
	val := ctx.Value(key)
	if val == nil {
		return zero, false
	}

	result, ok := val.(T)
	return result, ok
}

// GetAPIBaseURL returns the base URL for API endpoints
func GetAPIBaseURL(ctx context.Context, key string) string {
	var apiPrefix = "api"

	td, ok := GetCtx[*TemplGo](ctx, key)
	if !ok {
		log.Fatal("url-api not available")
	}

	if td.APIHost == "" {
		return ""
	}

	// If host includes protocol, return as is
	if len(td.APIHost) > 4 && (td.APIHost[:4] == "http" || td.APIHost[:5] == "https") {
		return td.APIHost
	}

	// Otherwise construct the URL with protocol
	protocol := "http" // Default to http for local development
	if td.APIPort == 443 {
		protocol = "https"
	}

	if td.APIPort == 80 || td.APIPort == 443 {
		return fmt.Sprintf("%s://%s/%s", protocol, td.APIHost, apiPrefix)
	}
	return fmt.Sprintf("%s://%s:%d/%s", protocol, td.APIHost, td.APIPort, apiPrefix)
}
