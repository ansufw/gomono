package main

import (
	"log"

	"github.com/ansufw/gomono/internal/apps"
)

func main() {
	app, err := apps.Wire()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	apps.Run(app)
}
