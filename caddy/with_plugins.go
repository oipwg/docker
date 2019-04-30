package main

import (
	"github.com/mholt/caddy/caddy/caddymain"

	// plug in plugins here
	_ "github.com/captncraig/cors/caddy"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
