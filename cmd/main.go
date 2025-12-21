package main

import (
	"github.com/ansufw/gomono/internal/apps"
)

func main() {
	gem := apps.Wire()
	apps.Run(gem)
}
