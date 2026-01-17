package checkbox

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

func TestCheckboxRendersInputTypeCheckbox(t *testing.T) {
	html := render(t, Checkbox(Props{
		Label: "Accept terms",
	}))

	if !strings.Contains(html, `type="checkbox"`) {
		t.Errorf("expected input type=checkbox, got: %s", html)
	}
	if !strings.Contains(html, "<input") {
		t.Errorf("expected input element, got: %s", html)
	}
}

func TestCheckboxLabelRendersAndAssociatesCorrectly(t *testing.T) {
	html := render(t, Checkbox(Props{
		Label: "Accept terms",
	}))

	// Label wraps the input (Pico CSS pattern)
	if !strings.Contains(html, "<label") {
		t.Errorf("expected label element, got: %s", html)
	}
	if !strings.Contains(html, "Accept terms") {
		t.Errorf("expected label text 'Accept terms', got: %s", html)
	}
	// The input should be inside the label
	labelStart := strings.Index(html, "<label")
	labelEnd := strings.Index(html, "</label>")
	inputPos := strings.Index(html, "<input")
	if inputPos < labelStart || inputPos > labelEnd {
		t.Errorf("expected input inside label element, got: %s", html)
	}
}

func TestCheckboxCheckedAddsCheckedAttribute(t *testing.T) {
	html := render(t, Checkbox(Props{
		Label:   "Remember me",
		Checked: true,
	}))

	if !strings.Contains(html, "checked") {
		t.Errorf("expected checked attribute, got: %s", html)
	}
}

func TestCheckboxUncheckedDoesNotHaveCheckedAttribute(t *testing.T) {
	html := render(t, Checkbox(Props{
		Label:   "Remember me",
		Checked: false,
	}))

	if strings.Contains(html, "checked") {
		t.Errorf("did not expect checked attribute, got: %s", html)
	}
}

func TestCheckboxDisabledAddsDisabledAttribute(t *testing.T) {
	html := render(t, Checkbox(Props{
		Label:    "Unavailable option",
		Disabled: true,
	}))

	if !strings.Contains(html, "disabled") {
		t.Errorf("expected disabled attribute on input, got: %s", html)
	}
	if !strings.Contains(html, `aria-disabled="true"`) {
		t.Errorf("expected aria-disabled on label, got: %s", html)
	}
}

func TestCheckboxValueAttributeRenders(t *testing.T) {
	html := render(t, Checkbox(Props{
		Label: "Option A",
		Value: "option_a",
	}))

	if !strings.Contains(html, `value="option_a"`) {
		t.Errorf("expected value attribute, got: %s", html)
	}
}

func TestCheckboxNameAndIDRender(t *testing.T) {
	html := render(t, Checkbox(Props{
		Name:  "terms",
		ID:    "terms-checkbox",
		Label: "I agree",
	}))

	if !strings.Contains(html, `name="terms"`) {
		t.Errorf("expected name attribute, got: %s", html)
	}
	if !strings.Contains(html, `id="terms-checkbox"`) {
		t.Errorf("expected id attribute, got: %s", html)
	}
}

func TestCheckboxInvalidAddsAriaInvalid(t *testing.T) {
	html := render(t, Checkbox(Props{
		Label:   "Required checkbox",
		Invalid: true,
	}))

	if !strings.Contains(html, `aria-invalid="true"`) {
		t.Errorf("expected aria-invalid attribute, got: %s", html)
	}
}

func TestCheckboxClassApplied(t *testing.T) {
	html := render(t, Checkbox(Props{
		Label: "Custom styled",
		Class: "custom-class",
	}))

	if !strings.Contains(html, `class="custom-class"`) {
		t.Errorf("expected class attribute on label, got: %s", html)
	}
}

func TestCheckboxAttrsSpread(t *testing.T) {
	html := render(t, Checkbox(Props{
		Label: "With custom attrs",
		Attrs: templ.Attributes{
			"data-testid": "my-checkbox",
		},
	}))

	if !strings.Contains(html, `data-testid="my-checkbox"`) {
		t.Errorf("expected custom attribute, got: %s", html)
	}
}
