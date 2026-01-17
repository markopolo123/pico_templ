package input

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func render(t *testing.T, c templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	if err := c.Render(context.Background(), &buf); err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	return buf.String()
}

func TestDefaultTextInput(t *testing.T) {
	html := render(t, Input(Props{
		Name: "username",
	}))

	if !strings.Contains(html, `type="text"`) {
		t.Error("expected type to default to text")
	}
	if !strings.Contains(html, `name="username"`) {
		t.Error("expected name attribute")
	}
	if !strings.Contains(html, `id="username"`) {
		t.Error("expected id to default to name")
	}
}

func TestLabelWrapsInput(t *testing.T) {
	html := render(t, Input(Props{
		Name:  "email",
		Label: "Email Address",
	}))

	if !strings.Contains(html, `<label`) {
		t.Error("expected label element when Label prop is set")
	}
	if !strings.Contains(html, `for="email"`) {
		t.Error("expected label to have for attribute matching input id")
	}
	if !strings.Contains(html, "Email Address") {
		t.Error("expected label text to be rendered")
	}
}

func TestPlaceholderRenders(t *testing.T) {
	html := render(t, Input(Props{
		Name:        "search",
		Placeholder: "Search...",
	}))

	if !strings.Contains(html, `placeholder="Search..."`) {
		t.Error("expected placeholder attribute")
	}
}

func TestRequiredAttribute(t *testing.T) {
	html := render(t, Input(Props{
		Name:     "password",
		Required: true,
	}))

	if !strings.Contains(html, "required") {
		t.Error("expected required attribute")
	}
}

func TestDisabledAttribute(t *testing.T) {
	html := render(t, Input(Props{
		Name:     "disabled-field",
		Disabled: true,
	}))

	if !strings.Contains(html, "disabled") {
		t.Error("expected disabled attribute")
	}
}

func TestInvalidAddsAriaInvalid(t *testing.T) {
	html := render(t, Input(Props{
		Name:    "invalid-field",
		Invalid: true,
	}))

	if !strings.Contains(html, `aria-invalid="true"`) {
		t.Error("expected aria-invalid=\"true\" attribute")
	}
}

func TestHelperTextRendersSmallElement(t *testing.T) {
	html := render(t, Input(Props{
		Name:       "field",
		HelperText: "This is helper text",
	}))

	if !strings.Contains(html, "<small") {
		t.Error("expected small element for helper text")
	}
	if !strings.Contains(html, "This is helper text") {
		t.Error("expected helper text content")
	}
	if !strings.Contains(html, `aria-describedby="field-helper"`) {
		t.Error("expected aria-describedby attribute linking to helper")
	}
	if !strings.Contains(html, `id="field-helper"`) {
		t.Error("expected helper text to have id")
	}
}

func TestInputTypes(t *testing.T) {
	types := []string{
		"text", "email", "password", "number", "tel", "url",
		"search", "date", "time", "datetime-local", "month", "week", "color",
	}

	for _, inputType := range types {
		t.Run(inputType, func(t *testing.T) {
			html := render(t, Input(Props{
				Name: "field",
				Type: inputType,
			}))

			expected := `type="` + inputType + `"`
			if !strings.Contains(html, expected) {
				t.Errorf("expected %s, got %s", expected, html)
			}
		})
	}
}

func TestIDDefaultsToName(t *testing.T) {
	html := render(t, Input(Props{
		Name: "myfield",
	}))

	if !strings.Contains(html, `id="myfield"`) {
		t.Error("expected id to default to Name when ID is not provided")
	}
}

func TestExplicitID(t *testing.T) {
	html := render(t, Input(Props{
		Name: "myfield",
		ID:   "custom-id",
	}))

	if !strings.Contains(html, `id="custom-id"`) {
		t.Error("expected explicit ID to be used")
	}
	if strings.Contains(html, `id="myfield"`) {
		t.Error("expected Name not to be used as id when ID is provided")
	}
}

func TestReadOnlyAttribute(t *testing.T) {
	html := render(t, Input(Props{
		Name:     "readonly-field",
		ReadOnly: true,
	}))

	if !strings.Contains(html, "readonly") {
		t.Error("expected readonly attribute")
	}
}

func TestValueAttribute(t *testing.T) {
	html := render(t, Input(Props{
		Name:  "field",
		Value: "prefilled value",
	}))

	if !strings.Contains(html, `value="prefilled value"`) {
		t.Error("expected value attribute")
	}
}

func TestClassAttribute(t *testing.T) {
	html := render(t, Input(Props{
		Name:  "field",
		Class: "custom-class",
	}))

	if !strings.Contains(html, `class="custom-class"`) {
		t.Error("expected class attribute")
	}
}

func TestCustomAttrs(t *testing.T) {
	html := render(t, Input(Props{
		Name: "field",
		Attrs: templ.Attributes{
			"data-custom": "value",
			"maxlength":   "100",
		},
	}))

	if !strings.Contains(html, `data-custom="value"`) {
		t.Error("expected custom data attribute")
	}
	if !strings.Contains(html, `maxlength="100"`) {
		t.Error("expected maxlength attribute")
	}
}
