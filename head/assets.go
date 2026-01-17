package head

import "embed"

//go:embed assets/*.css assets/*.js
var Assets embed.FS

// GetPicoCSS returns the embedded Pico CSS content.
func GetPicoCSS() ([]byte, error) {
	return Assets.ReadFile("assets/pico.min.css")
}

// GetHTMX returns the embedded HTMX JavaScript content.
func GetHTMX() ([]byte, error) {
	return Assets.ReadFile("assets/htmx.min.js")
}

// GetHyperscript returns the embedded _hyperscript JavaScript content.
func GetHyperscript() ([]byte, error) {
	return Assets.ReadFile("assets/_hyperscript.min.js")
}
