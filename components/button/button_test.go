package button

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func render(t *testing.T, component templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	err := component.Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("failed to render component: %v", err)
	}
	return buf.String()
}

func TestDefaultButton(t *testing.T) {
	html := render(t, Button(Props{Text: "Click me"}))
	if !strings.Contains(html, `<button type="button"`) {
		t.Errorf("expected button with type=button, got: %s", html)
	}
	if !strings.Contains(html, `>Click me</button>`) {
		t.Errorf("expected button text 'Click me', got: %s", html)
	}
}

func TestSecondaryVariant(t *testing.T) {
	html := render(t, Button(Props{Text: "Secondary", Variant: Secondary}))
	if !strings.Contains(html, `class="secondary"`) {
		t.Errorf("expected .secondary class, got: %s", html)
	}
}

func TestContrastVariant(t *testing.T) {
	html := render(t, Button(Props{Text: "Contrast", Variant: Contrast}))
	if !strings.Contains(html, `class="contrast"`) {
		t.Errorf("expected .contrast class, got: %s", html)
	}
}

func TestOutline(t *testing.T) {
	html := render(t, Button(Props{Text: "Outline", Outline: true}))
	if !strings.Contains(html, `class="outline"`) {
		t.Errorf("expected .outline class, got: %s", html)
	}
}

func TestOutlineSecondary(t *testing.T) {
	html := render(t, Button(Props{Text: "Outline Secondary", Outline: true, Variant: Secondary}))
	if !strings.Contains(html, `outline`) || !strings.Contains(html, `secondary`) {
		t.Errorf("expected .outline.secondary classes, got: %s", html)
	}
}

func TestDisabled(t *testing.T) {
	html := render(t, Button(Props{Text: "Disabled", Disabled: true}))
	if !strings.Contains(html, `disabled`) {
		t.Errorf("expected disabled attribute, got: %s", html)
	}
}

func TestHtmxAttributes(t *testing.T) {
	html := render(t, Button(Props{
		Text:      "HTMX",
		HxGet:     "/api/data",
		HxPost:    "/api/submit",
		HxPut:     "/api/update",
		HxDelete:  "/api/delete",
		HxPatch:   "/api/patch",
		HxTarget:  "#result",
		HxSwap:    "innerHTML",
		HxTrigger: "click",
	}))
	expectations := []string{
		`hx-get="/api/data"`,
		`hx-post="/api/submit"`,
		`hx-put="/api/update"`,
		`hx-delete="/api/delete"`,
		`hx-patch="/api/patch"`,
		`hx-target="#result"`,
		`hx-swap="innerHTML"`,
		`hx-trigger="click"`,
	}
	for _, exp := range expectations {
		if !strings.Contains(html, exp) {
			t.Errorf("expected %s, got: %s", exp, html)
		}
	}
}

func TestTypeSubmit(t *testing.T) {
	html := render(t, Button(Props{Text: "Submit", Type: "submit"}))
	if !strings.Contains(html, `type="submit"`) {
		t.Errorf("expected type=submit, got: %s", html)
	}
}

func TestTypeReset(t *testing.T) {
	html := render(t, Button(Props{Text: "Reset", Type: "reset"}))
	if !strings.Contains(html, `type="reset"`) {
		t.Errorf("expected type=reset, got: %s", html)
	}
}

func TestCustomClass(t *testing.T) {
	html := render(t, Button(Props{Text: "Custom", Variant: Secondary, Class: "my-class"}))
	if !strings.Contains(html, `secondary`) || !strings.Contains(html, `my-class`) {
		t.Errorf("expected secondary and my-class classes, got: %s", html)
	}
}

func TestAttrsSpread(t *testing.T) {
	html := render(t, Button(Props{
		Text: "Custom Attrs",
		Attrs: templ.Attributes{
			"data-testid": "my-button",
			"aria-label":  "Custom button",
		},
	}))
	if !strings.Contains(html, `data-testid="my-button"`) {
		t.Errorf("expected data-testid attribute, got: %s", html)
	}
	if !strings.Contains(html, `aria-label="Custom button"`) {
		t.Errorf("expected aria-label attribute, got: %s", html)
	}
}
