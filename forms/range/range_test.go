package rangecomp

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func renderToString(t *testing.T, component templ.Component) string {
	t.Helper()
	var buf strings.Builder
	err := component.Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("failed to render component: %v", err)
	}
	return buf.String()
}

func TestRangeRendersInputTypeRange(t *testing.T) {
	html := renderToString(t, Range(Props{}))

	if !strings.Contains(html, `type="range"`) {
		t.Error("expected input to have type=\"range\"")
	}
	if !strings.Contains(html, "<input") {
		t.Error("expected an input element")
	}
}

func TestRangeMinMaxStepAttributes(t *testing.T) {
	html := renderToString(t, Range(Props{
		Min:  0,
		Max:  100,
		Step: 5,
	}))

	if !strings.Contains(html, `min="0"`) {
		t.Errorf("expected min=\"0\", got: %s", html)
	}
	if !strings.Contains(html, `max="100"`) {
		t.Errorf("expected max=\"100\", got: %s", html)
	}
	if !strings.Contains(html, `step="5"`) {
		t.Errorf("expected step=\"5\", got: %s", html)
	}
}

func TestRangeValueSetsInitialPosition(t *testing.T) {
	html := renderToString(t, Range(Props{
		Value: 50,
	}))

	if !strings.Contains(html, `value="50"`) {
		t.Errorf("expected value=\"50\", got: %s", html)
	}
}

func TestRangeLabelAssociatesCorrectly(t *testing.T) {
	html := renderToString(t, Range(Props{
		Label: "Volume",
		ID:    "volume-slider",
	}))

	if !strings.Contains(html, "<label>") {
		t.Error("expected a label element")
	}
	if !strings.Contains(html, "Volume") {
		t.Error("expected label text 'Volume'")
	}
	if !strings.Contains(html, `id="volume-slider"`) {
		t.Errorf("expected id=\"volume-slider\", got: %s", html)
	}
}

func TestRangeDisabledAttribute(t *testing.T) {
	html := renderToString(t, Range(Props{
		Disabled: true,
	}))

	if !strings.Contains(html, "disabled") {
		t.Error("expected disabled attribute")
	}
}

func TestRangeNameAttribute(t *testing.T) {
	html := renderToString(t, Range(Props{
		Name: "brightness",
	}))

	if !strings.Contains(html, `name="brightness"`) {
		t.Errorf("expected name=\"brightness\", got: %s", html)
	}
}

func TestRangeClassAttribute(t *testing.T) {
	html := renderToString(t, Range(Props{
		Class: "custom-class",
	}))

	if !strings.Contains(html, `class="custom-class"`) {
		t.Errorf("expected class=\"custom-class\", got: %s", html)
	}
}

func TestRangeCustomAttrs(t *testing.T) {
	html := renderToString(t, Range(Props{
		Attrs: templ.Attributes{
			"data-testid": "my-range",
			"aria-label":  "Volume control",
		},
	}))

	if !strings.Contains(html, `data-testid="my-range"`) {
		t.Errorf("expected data-testid attribute, got: %s", html)
	}
	if !strings.Contains(html, `aria-label="Volume control"`) {
		t.Errorf("expected aria-label attribute, got: %s", html)
	}
}

func TestRangeFloatValues(t *testing.T) {
	html := renderToString(t, Range(Props{
		Min:   0.5,
		Max:   10.5,
		Step:  0.1,
		Value: 5.5,
	}))

	if !strings.Contains(html, `min="0.5"`) {
		t.Errorf("expected min=\"0.5\", got: %s", html)
	}
	if !strings.Contains(html, `max="10.5"`) {
		t.Errorf("expected max=\"10.5\", got: %s", html)
	}
	if !strings.Contains(html, `step="0.1"`) {
		t.Errorf("expected step=\"0.1\", got: %s", html)
	}
	if !strings.Contains(html, `value="5.5"`) {
		t.Errorf("expected value=\"5.5\", got: %s", html)
	}
}
