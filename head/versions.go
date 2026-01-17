// Package head provides a Head component with embedded versioned assets.
package head

const (
	PicoCSSVersion     = "2.1.1"
	HTMXVersion        = "2.0.4"
	HyperscriptVersion = "0.9.14"
)

// AssetHashes contains expected SHA256 hashes for integrity verification.
var AssetHashes = map[string]string{
	"pico.min.css":        "d909404e60ea5ddec11a48b55292f110f713c6c30ab4d9b9bfaa0f31f363ca6f",
	"htmx.min.js":         "69caefd0da92269066e725d7fe175e26b9d50c962e3056459c0c477154cdb9d3",
	"_hyperscript.min.js": "3e834a3ffc0334fee54ecff4e37a6ae951cd83e6daa96651ca7cfd8f751ad4d2",
}

// AssetURLs contains download URLs for each asset.
var AssetURLs = map[string]string{
	"pico.min.css":        "https://cdn.jsdelivr.net/npm/@picocss/pico@" + PicoCSSVersion + "/css/pico.min.css",
	"htmx.min.js":         "https://unpkg.com/htmx.org@" + HTMXVersion + "/dist/htmx.min.js",
	"_hyperscript.min.js": "https://unpkg.com/hyperscript.org@" + HyperscriptVersion + "/dist/_hyperscript.min.js",
}
