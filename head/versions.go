// Package head provides a Head component with embedded versioned assets.
package head

const (
	PicoCSSVersion     = "2.1.1"
	HTMXVersion        = "2.0.4"
	HyperscriptVersion = "0.9.14"
)

// AssetHashes contains expected SHA256 hashes for integrity verification.
var AssetHashes = map[string]string{
	"pico.min.css":        "fbc9a63fc9fc9f72d12fd7fc9806e11fa9f77ae4f9cad146b27003a1119ba3db",
	"htmx.min.js":         "e209dda5c8235479f3166defc7750e1dbcd5a5c1808b7792fc2e6733768fb447",
	"_hyperscript.min.js": "3e834a3ffc0334fee54ecff4e37a6ae951cd83e6daa96651ca7cfd8f751ad4d2",
}

// AssetURLs contains download URLs for each asset.
var AssetURLs = map[string]string{
	"pico.min.css":        "https://cdn.jsdelivr.net/npm/@picocss/pico@" + PicoCSSVersion + "/css/pico.min.css",
	"htmx.min.js":         "https://unpkg.com/htmx.org@" + HTMXVersion + "/dist/htmx.min.js",
	"_hyperscript.min.js": "https://unpkg.com/hyperscript.org@" + HyperscriptVersion + "/dist/_hyperscript.min.js",
}
