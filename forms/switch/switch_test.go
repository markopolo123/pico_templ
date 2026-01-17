package switch_

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func render(t *testing.T, component templ.Component) string {
	t.Helper()
	var buf strings.Builder
	err := component.Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("failed to render component: %v", err)
	}
	return buf.String()
}

func TestSwitchRendersCheckboxWithRoleSwitch(t *testing.T) {
	html := render(t, Switch(Props{
		Label: "Toggle me",
	}))

	if !strings.Contains(html, `type="checkbox"`) {
		t.Error("expected input type='checkbox'")
	}
	if !strings.Contains(html, `role="switch"`) {
		t.Error("expected role='switch'")
	}
}

func TestSwitchLabelWrapsInput(t *testing.T) {
	html := render(t, Switch(Props{
		Label: "My Label",
	}))

	// Label should contain input and text
	if !strings.Contains(html, "<label") {
		t.Error("expected label element")
	}
	if !strings.Contains(html, "My Label") {
		t.Error("expected label text")
	}
	// Input should be inside label (label comes before input in HTML)
	labelIdx := strings.Index(html, "<label")
	inputIdx := strings.Index(html, "<input")
	closeLabelIdx := strings.Index(html, "</label>")
	if labelIdx == -1 || inputIdx == -1 || closeLabelIdx == -1 {
		t.Error("expected label to wrap input")
	}
	if !(labelIdx < inputIdx && inputIdx < closeLabelIdx) {
		t.Error("expected input to be inside label")
	}
}

func TestSwitchCheckedState(t *testing.T) {
	// Test unchecked (default)
	html := render(t, Switch(Props{
		Label:   "Off",
		Checked: false,
	}))
	if strings.Contains(html, " checked") {
		t.Error("expected no checked attribute when Checked=false")
	}

	// Test checked
	html = render(t, Switch(Props{
		Label:   "On",
		Checked: true,
	}))
	if !strings.Contains(html, " checked") {
		t.Error("expected checked attribute when Checked=true")
	}
}

func TestSwitchDisabledState(t *testing.T) {
	// Test enabled (default)
	html := render(t, Switch(Props{
		Label:    "Active",
		Disabled: false,
	}))
	if strings.Contains(html, " disabled") {
		t.Error("expected no disabled attribute when Disabled=false")
	}

	// Test disabled
	html = render(t, Switch(Props{
		Label:    "Inactive",
		Disabled: true,
	}))
	if !strings.Contains(html, " disabled") {
		t.Error("expected disabled attribute when Disabled=true")
	}
}

func TestSwitchNameAndID(t *testing.T) {
	html := render(t, Switch(Props{
		Name:  "my-switch",
		ID:    "switch-id",
		Label: "Named Switch",
	}))

	if !strings.Contains(html, `name="my-switch"`) {
		t.Error("expected name attribute")
	}
	if !strings.Contains(html, `id="switch-id"`) {
		t.Error("expected id attribute")
	}
}

func TestSwitchCustomClass(t *testing.T) {
	html := render(t, Switch(Props{
		Label: "Styled",
		Class: "custom-class",
	}))

	if !strings.Contains(html, `class="custom-class"`) {
		t.Error("expected custom class on label")
	}
}

func TestSwitchCustomAttrs(t *testing.T) {
	html := render(t, Switch(Props{
		Label: "With attrs",
		Attrs: templ.Attributes{
			"data-testid": "my-switch",
			"aria-label":  "Toggle feature",
		},
	}))

	if !strings.Contains(html, `data-testid="my-switch"`) {
		t.Error("expected data-testid attribute")
	}
	if !strings.Contains(html, `aria-label="Toggle feature"`) {
		t.Error("expected aria-label attribute")
	}
}
