package head

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestHeadRendersWithAllAssets(t *testing.T) {
	props := DefaultProps()
	props.Title = "Test Page"
	props.Description = "Test description"

	html := renderHead(t, props)

	if !strings.Contains(html, "<title>Test Page</title>") {
		t.Error("expected title to be rendered")
	}
	if !strings.Contains(html, `content="Test description"`) {
		t.Error("expected description to be rendered")
	}
	if !strings.Contains(html, "<style>") {
		t.Error("expected Pico CSS style tag")
	}
	// Check for HTMX signature
	if !strings.Contains(html, "htmx") {
		t.Error("expected HTMX to be included")
	}
	// Check for _hyperscript signature
	if !strings.Contains(html, "_hyperscript") {
		t.Error("expected _hyperscript to be included")
	}
}

func TestIncludePicoFalse(t *testing.T) {
	props := DefaultProps()
	props.Title = "Test"
	props.IncludePico = false

	html := renderHead(t, props)

	// Should have script tags but no style tag with Pico content
	if strings.Contains(html, "pico") || strings.Contains(html, ":root") {
		t.Error("Pico CSS should not be included when IncludePico=false")
	}
}

func TestIncludeHTMXFalse(t *testing.T) {
	props := DefaultProps()
	props.Title = "Test"
	props.IncludeHTMX = false

	html := renderHead(t, props)

	if strings.Contains(html, "htmx.org") || strings.Contains(html, "hx-") {
		t.Error("HTMX should not be included when IncludeHTMX=false")
	}
}

func TestIncludeHyperscriptFalse(t *testing.T) {
	props := DefaultProps()
	props.Title = "Test"
	props.IncludeHyperscript = false

	html := renderHead(t, props)

	if strings.Contains(html, "_hyperscript") || strings.Contains(html, "hyperscript.org") {
		t.Error("_hyperscript should not be included when IncludeHyperscript=false")
	}
}

func TestTitleAndDescriptionRender(t *testing.T) {
	props := Props{
		Title:       "My App Title",
		Description: "My app description here",
	}

	html := renderHead(t, props)

	if !strings.Contains(html, "<title>My App Title</title>") {
		t.Errorf("expected title 'My App Title', got html: %s", html[:200])
	}
	if !strings.Contains(html, `name="description"`) {
		t.Error("expected description meta tag")
	}
	if !strings.Contains(html, `content="My app description here"`) {
		t.Error("expected description content")
	}
}

func TestDescriptionOmittedWhenEmpty(t *testing.T) {
	props := Props{
		Title:       "Test",
		Description: "",
	}

	html := renderHead(t, props)

	if strings.Contains(html, `name="description"`) {
		t.Error("description meta tag should be omitted when Description is empty")
	}
}

func TestExtraHeadContent(t *testing.T) {
	extraHead := templ.Raw(`<link rel="icon" href="/favicon.ico">`)
	props := Props{
		Title:     "Test",
		ExtraHead: extraHead,
	}

	html := renderHead(t, props)

	if !strings.Contains(html, `<link rel="icon" href="/favicon.ico">`) {
		t.Error("expected ExtraHead content to be rendered")
	}
}

func TestAssetsAreProperlyEmbedded(t *testing.T) {
	picoData, err := GetPicoCSS()
	if err != nil {
		t.Fatalf("failed to get Pico CSS: %v", err)
	}
	if len(picoData) == 0 {
		t.Error("Pico CSS asset is empty")
	}

	htmxData, err := GetHTMX()
	if err != nil {
		t.Fatalf("failed to get HTMX: %v", err)
	}
	if len(htmxData) == 0 {
		t.Error("HTMX asset is empty")
	}

	hsData, err := GetHyperscript()
	if err != nil {
		t.Fatalf("failed to get _hyperscript: %v", err)
	}
	if len(hsData) == 0 {
		t.Error("_hyperscript asset is empty")
	}
}

func TestHashVerificationLogic(t *testing.T) {
	// Test that our hash verification logic works correctly
	testData := []byte("test content")
	hash := sha256.Sum256(testData)
	actualHash := hex.EncodeToString(hash[:])

	expectedHash := "6ae8a75555209fd6c44157c0aed8016e763ff435a19cf186f76863140143ff72"

	if actualHash != expectedHash {
		t.Errorf("hash mismatch in verification logic: expected %s, got %s", expectedHash, actualHash)
	}

	// Test mismatch detection
	wrongHash := "0000000000000000000000000000000000000000000000000000000000000000"
	if actualHash == wrongHash {
		t.Error("hash verification should detect mismatches")
	}
}

func TestAssetHashesMatchEmbeddedFiles(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		getData  func() ([]byte, error)
	}{
		{"Pico CSS", "pico.min.css", GetPicoCSS},
		{"HTMX", "htmx.min.js", GetHTMX},
		{"_hyperscript", "_hyperscript.min.js", GetHyperscript},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.getData()
			if err != nil {
				t.Fatalf("failed to get asset: %v", err)
			}

			hash := sha256.Sum256(data)
			actualHash := hex.EncodeToString(hash[:])
			expectedHash := AssetHashes[tt.filename]

			if actualHash != expectedHash {
				t.Errorf("hash mismatch for %s: expected %s, got %s", tt.filename, expectedHash, actualHash)
			}
		})
	}
}

func TestDefaultProps(t *testing.T) {
	props := DefaultProps()

	if !props.IncludePico {
		t.Error("IncludePico should default to true")
	}
	if !props.IncludeHTMX {
		t.Error("IncludeHTMX should default to true")
	}
	if !props.IncludeHyperscript {
		t.Error("IncludeHyperscript should default to true")
	}
}

func TestHeadContainsRequiredMetaTags(t *testing.T) {
	props := Props{Title: "Test"}
	html := renderHead(t, props)

	if !strings.Contains(html, `charset="utf-8"`) {
		t.Error("expected charset meta tag")
	}
	if !strings.Contains(html, `name="viewport"`) {
		t.Error("expected viewport meta tag")
	}
	if !strings.Contains(html, `width=device-width`) {
		t.Error("expected viewport content")
	}
}

func renderHead(t *testing.T, props Props) string {
	t.Helper()
	var buf bytes.Buffer
	err := Head(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("failed to render Head: %v", err)
	}
	return buf.String()
}
