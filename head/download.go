//go:build ignore

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	PicoCSSVersion     = "2.1.1"
	HTMXVersion        = "2.0.4"
	HyperscriptVersion = "0.9.14"
)

var assetHashes = map[string]string{
	"pico.min.css":        "fbc9a63fc9fc9f72d12fd7fc9806e11fa9f77ae4f9cad146b27003a1119ba3db",
	"htmx.min.js":         "e209dda5c8235479f3166defc7750e1dbcd5a5c1808b7792fc2e6733768fb447",
	"_hyperscript.min.js": "3e834a3ffc0334fee54ecff4e37a6ae951cd83e6daa96651ca7cfd8f751ad4d2",
}

var assetURLs = map[string]string{
	"pico.min.css":        "https://cdn.jsdelivr.net/npm/@picocss/pico@" + PicoCSSVersion + "/css/pico.min.css",
	"htmx.min.js":         "https://unpkg.com/htmx.org@" + HTMXVersion + "/dist/htmx.min.js",
	"_hyperscript.min.js": "https://unpkg.com/hyperscript.org@" + HyperscriptVersion + "/dist/_hyperscript.min.js",
}

func main() {
	assetsDir := "assets"
	if err := os.MkdirAll(assetsDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create assets directory: %v\n", err)
		os.Exit(1)
	}

	for filename, url := range assetURLs {
		expectedHash := assetHashes[filename]
		destPath := filepath.Join(assetsDir, filename)

		fmt.Printf("Downloading %s...\n", filename)

		if err := downloadAndVerify(url, destPath, expectedHash); err != nil {
			fmt.Fprintf(os.Stderr, "failed to download %s: %v\n", filename, err)
			os.Exit(1)
		}

		fmt.Printf("  ✓ %s verified and saved\n", filename)
	}

	fmt.Println("All assets downloaded and verified successfully")
}

func downloadAndVerify(url, destPath, expectedHash string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("HTTP GET failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	hash := sha256.Sum256(data)
	actualHash := hex.EncodeToString(hash[:])

	if actualHash != expectedHash {
		return fmt.Errorf("hash mismatch: expected %s, got %s", expectedHash, actualHash)
	}

	if err := os.WriteFile(destPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
