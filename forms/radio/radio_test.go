package radio

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestRadioRendersInputTypeRadio(t *testing.T) {
	props := Props{
		Name:  "test",
		Label: "Test Option",
	}

	var buf bytes.Buffer
	err := Radio(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html := buf.String()
	if !strings.Contains(html, `type="radio"`) {
		t.Errorf("expected input type='radio', got: %s", html)
	}
}

func TestRadioNameGroupsRadiosTogether(t *testing.T) {
	props := Props{
		Name:  "language",
		Label: "English",
	}

	var buf bytes.Buffer
	err := Radio(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html := buf.String()
	if !strings.Contains(html, `name="language"`) {
		t.Errorf("expected name='language', got: %s", html)
	}
}

func TestRadioCheckedAddsCheckedAttribute(t *testing.T) {
	props := Props{
		Name:    "test",
		Label:   "Test Option",
		Checked: true,
	}

	var buf bytes.Buffer
	err := Radio(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html := buf.String()
	if !strings.Contains(html, "checked") {
		t.Errorf("expected checked attribute, got: %s", html)
	}
}

func TestRadioLabelAssociatesCorrectly(t *testing.T) {
	props := Props{
		Name:  "test",
		Label: "My Label Text",
	}

	var buf bytes.Buffer
	err := Radio(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html := buf.String()
	// Label should wrap the input and contain the label text
	if !strings.Contains(html, "<label") {
		t.Errorf("expected label element, got: %s", html)
	}
	if !strings.Contains(html, "My Label Text") {
		t.Errorf("expected label text 'My Label Text', got: %s", html)
	}
}

func TestRadioGroupRendersMultipleOptions(t *testing.T) {
	props := GroupProps{
		Name: "language",
		Options: []Props{
			{Label: "English", Value: "en", Checked: true},
			{Label: "French", Value: "fr"},
			{Label: "Spanish", Value: "es"},
		},
	}

	var buf bytes.Buffer
	err := RadioGroup(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html := buf.String()

	// Should render fieldset
	if !strings.Contains(html, "<fieldset") {
		t.Errorf("expected fieldset element, got: %s", html)
	}

	// Should render all options with shared name
	if strings.Count(html, `name="language"`) != 3 {
		t.Errorf("expected 3 radio inputs with name='language', got: %s", html)
	}

	// Should render all labels
	if !strings.Contains(html, "English") {
		t.Errorf("expected 'English' label, got: %s", html)
	}
	if !strings.Contains(html, "French") {
		t.Errorf("expected 'French' label, got: %s", html)
	}
	if !strings.Contains(html, "Spanish") {
		t.Errorf("expected 'Spanish' label, got: %s", html)
	}
}

func TestRadioDisabledAddsAttributes(t *testing.T) {
	props := Props{
		Name:     "test",
		Label:    "Disabled Option",
		Disabled: true,
	}

	var buf bytes.Buffer
	err := Radio(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html := buf.String()
	if !strings.Contains(html, "disabled") {
		t.Errorf("expected disabled attribute on input, got: %s", html)
	}
	if !strings.Contains(html, `aria-disabled="true"`) {
		t.Errorf("expected aria-disabled on label, got: %s", html)
	}
}

func TestRadioInvalidAddsAriaInvalid(t *testing.T) {
	props := Props{
		Name:    "test",
		Label:   "Invalid Option",
		Invalid: true,
	}

	var buf bytes.Buffer
	err := Radio(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html := buf.String()
	if !strings.Contains(html, `aria-invalid="true"`) {
		t.Errorf("expected aria-invalid='true', got: %s", html)
	}
}

func TestRadioValueAttribute(t *testing.T) {
	props := Props{
		Name:  "test",
		Label: "Test",
		Value: "test-value",
	}

	var buf bytes.Buffer
	err := Radio(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html := buf.String()
	if !strings.Contains(html, `value="test-value"`) {
		t.Errorf("expected value='test-value', got: %s", html)
	}
}

func TestRadioIDAttribute(t *testing.T) {
	props := Props{
		Name:  "test",
		ID:    "my-radio-id",
		Label: "Test",
	}

	var buf bytes.Buffer
	err := Radio(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html := buf.String()
	if !strings.Contains(html, `id="my-radio-id"`) {
		t.Errorf("expected id='my-radio-id', got: %s", html)
	}
}
