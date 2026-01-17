package selectfield

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func render(t *testing.T, props Props) string {
	t.Helper()
	var buf bytes.Buffer
	err := Select(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	return buf.String()
}

func TestSelectRendersWithOptions(t *testing.T) {
	html := render(t, Props{
		Name: "color",
		ID:   "color-select",
		Options: []Option{
			{Value: "red", Label: "Red"},
			{Value: "green", Label: "Green"},
			{Value: "blue", Label: "Blue"},
		},
	})

	if !strings.Contains(html, `<select`) {
		t.Error("expected select element")
	}
	if !strings.Contains(html, `name="color"`) {
		t.Error("expected name attribute")
	}
	if !strings.Contains(html, `id="color-select"`) {
		t.Error("expected id attribute")
	}
	if !strings.Contains(html, `value="red"`) {
		t.Error("expected red option value")
	}
	if !strings.Contains(html, `>Red</option>`) {
		t.Error("expected Red option label")
	}
	if !strings.Contains(html, `value="green"`) {
		t.Error("expected green option value")
	}
	if !strings.Contains(html, `value="blue"`) {
		t.Error("expected blue option value")
	}
}

func TestPlaceholderRendersAsDisabledFirstOption(t *testing.T) {
	html := render(t, Props{
		Name:        "fruit",
		Placeholder: "Choose a fruit...",
		Options: []Option{
			{Value: "apple", Label: "Apple"},
		},
	})

	if !strings.Contains(html, `<option selected disabled value="">Choose a fruit...</option>`) {
		t.Errorf("expected placeholder as disabled first option, got: %s", html)
	}

	// Verify placeholder comes before other options
	placeholderIdx := strings.Index(html, "Choose a fruit...")
	appleIdx := strings.Index(html, "Apple")
	if placeholderIdx > appleIdx {
		t.Error("placeholder should appear before other options")
	}
}

func TestOptGroupsRenderCorrectly(t *testing.T) {
	html := render(t, Props{
		Name: "car",
		OptGroups: []OptGroup{
			{
				Label: "Swedish Cars",
				Options: []Option{
					{Value: "volvo", Label: "Volvo"},
					{Value: "saab", Label: "Saab"},
				},
			},
			{
				Label: "German Cars",
				Options: []Option{
					{Value: "mercedes", Label: "Mercedes"},
					{Value: "audi", Label: "Audi"},
				},
			},
		},
	})

	if !strings.Contains(html, `<optgroup label="Swedish Cars">`) {
		t.Error("expected Swedish Cars optgroup")
	}
	if !strings.Contains(html, `<optgroup label="German Cars">`) {
		t.Error("expected German Cars optgroup")
	}
	if !strings.Contains(html, `value="volvo"`) {
		t.Error("expected volvo option")
	}
	if !strings.Contains(html, `value="mercedes"`) {
		t.Error("expected mercedes option")
	}
}

func TestSelectedOptionHasSelectedAttribute(t *testing.T) {
	html := render(t, Props{
		Name: "size",
		Options: []Option{
			{Value: "small", Label: "Small"},
			{Value: "medium", Label: "Medium", Selected: true},
			{Value: "large", Label: "Large"},
		},
	})

	// The selected option should have the selected attribute
	if !strings.Contains(html, `<option value="medium" selected>Medium</option>`) {
		t.Errorf("expected medium option to be selected, got: %s", html)
	}

	// Other options should not have selected attribute
	if strings.Contains(html, `value="small" selected`) {
		t.Error("small should not be selected")
	}
	if strings.Contains(html, `value="large" selected`) {
		t.Error("large should not be selected")
	}
}

func TestValidationStates(t *testing.T) {
	// Test invalid state
	html := render(t, Props{
		Name:    "required-field",
		Invalid: true,
		Options: []Option{
			{Value: "opt1", Label: "Option 1"},
		},
	})

	if !strings.Contains(html, `aria-invalid="true"`) {
		t.Error("expected aria-invalid=true for invalid state")
	}

	// Test valid state (no Invalid flag)
	htmlValid := render(t, Props{
		Name: "valid-field",
		Options: []Option{
			{Value: "opt1", Label: "Option 1"},
		},
	})

	if strings.Contains(htmlValid, `aria-invalid`) {
		t.Error("expected no aria-invalid attribute for valid state")
	}
}

func TestLabelRendersCorrectly(t *testing.T) {
	html := render(t, Props{
		Name:  "country",
		ID:    "country-select",
		Label: "Select Country",
		Options: []Option{
			{Value: "us", Label: "United States"},
		},
	})

	if !strings.Contains(html, `<label for="country-select">Select Country</label>`) {
		t.Errorf("expected label element, got: %s", html)
	}
}

func TestHelperTextRendersCorrectly(t *testing.T) {
	html := render(t, Props{
		Name:       "priority",
		HelperText: "Choose the priority level",
		Options: []Option{
			{Value: "low", Label: "Low"},
		},
	})

	if !strings.Contains(html, `<small>Choose the priority level</small>`) {
		t.Errorf("expected helper text, got: %s", html)
	}
}

func TestDisabledState(t *testing.T) {
	html := render(t, Props{
		Name:     "disabled-select",
		Disabled: true,
		Options: []Option{
			{Value: "opt1", Label: "Option 1"},
		},
	})

	if !strings.Contains(html, `disabled`) {
		t.Error("expected disabled attribute on select")
	}
}

func TestRequiredAttribute(t *testing.T) {
	html := render(t, Props{
		Name:     "required-select",
		Required: true,
		Options: []Option{
			{Value: "opt1", Label: "Option 1"},
		},
	})

	if !strings.Contains(html, `required`) {
		t.Error("expected required attribute on select")
	}
}

func TestDisabledOption(t *testing.T) {
	html := render(t, Props{
		Name: "option-disabled",
		Options: []Option{
			{Value: "available", Label: "Available"},
			{Value: "unavailable", Label: "Unavailable", Disabled: true},
		},
	})

	if !strings.Contains(html, `<option value="unavailable" disabled>Unavailable</option>`) {
		t.Errorf("expected disabled option, got: %s", html)
	}
}
