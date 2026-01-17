package textarea

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func render(t *testing.T, props Props) string {
	t.Helper()
	var buf bytes.Buffer
	err := Textarea(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	return buf.String()
}

func TestDefaultTextarea(t *testing.T) {
	html := render(t, Props{
		Name: "bio",
	})

	if !strings.Contains(html, `<textarea`) {
		t.Error("expected textarea element")
	}
	if !strings.Contains(html, `name="bio"`) {
		t.Error("expected name attribute")
	}
	if !strings.Contains(html, `id="bio"`) {
		t.Error("expected id to default to name")
	}
}

func TestLabelWrapsTextarea(t *testing.T) {
	html := render(t, Props{
		Name:  "bio",
		Label: "Biography",
	})

	if !strings.Contains(html, `<label`) {
		t.Error("expected label element")
	}
	if !strings.Contains(html, `for="bio"`) {
		t.Error("expected for attribute on label")
	}
	if !strings.Contains(html, `Biography`) {
		t.Error("expected label text")
	}
	// Label should wrap textarea
	labelIdx := strings.Index(html, "<label")
	textareaIdx := strings.Index(html, "<textarea")
	closeLabelIdx := strings.Index(html, "</label>")
	if labelIdx == -1 || textareaIdx == -1 || closeLabelIdx == -1 {
		t.Fatal("missing expected elements")
	}
	if !(labelIdx < textareaIdx && textareaIdx < closeLabelIdx) {
		t.Error("expected label to wrap textarea")
	}
}

func TestRowsAttribute(t *testing.T) {
	html := render(t, Props{
		Name: "bio",
		Rows: 5,
	})

	if !strings.Contains(html, `rows="5"`) {
		t.Error("expected rows attribute")
	}
}

func TestRowsNotRenderedWhenZero(t *testing.T) {
	html := render(t, Props{
		Name: "bio",
		Rows: 0,
	})

	if strings.Contains(html, `rows=`) {
		t.Error("rows should not render when zero")
	}
}

func TestValueRendersAsContent(t *testing.T) {
	html := render(t, Props{
		Name:  "bio",
		Value: "Hello World",
	})

	if !strings.Contains(html, `>Hello World</textarea>`) {
		t.Error("expected value as textarea content")
	}
}

func TestInvalidState(t *testing.T) {
	html := render(t, Props{
		Name:    "bio",
		Invalid: true,
	})

	if !strings.Contains(html, `aria-invalid="true"`) {
		t.Error("expected aria-invalid attribute")
	}
}

func TestHelperTextRenders(t *testing.T) {
	html := render(t, Props{
		Name:       "bio",
		HelperText: "Enter your biography",
	})

	if !strings.Contains(html, `<small`) {
		t.Error("expected small element for helper text")
	}
	if !strings.Contains(html, `Enter your biography`) {
		t.Error("expected helper text content")
	}
	if !strings.Contains(html, `id="bio-helper"`) {
		t.Error("expected helper id")
	}
	if !strings.Contains(html, `aria-describedby="bio-helper"`) {
		t.Error("expected aria-describedby on textarea")
	}
}

func TestHelperTextWithLabel(t *testing.T) {
	html := render(t, Props{
		Name:       "bio",
		Label:      "Biography",
		HelperText: "Enter your biography",
	})

	// Helper text should be inside label
	labelIdx := strings.Index(html, "<label")
	smallIdx := strings.Index(html, "<small")
	closeLabelIdx := strings.Index(html, "</label>")

	if labelIdx == -1 || smallIdx == -1 || closeLabelIdx == -1 {
		t.Fatal("missing expected elements")
	}
	if !(labelIdx < smallIdx && smallIdx < closeLabelIdx) {
		t.Error("expected helper text inside label")
	}
}

func TestDisabledState(t *testing.T) {
	html := render(t, Props{
		Name:     "bio",
		Disabled: true,
	})

	if !strings.Contains(html, ` disabled`) {
		t.Error("expected disabled attribute")
	}
}

func TestReadOnlyState(t *testing.T) {
	html := render(t, Props{
		Name:     "bio",
		ReadOnly: true,
	})

	if !strings.Contains(html, ` readonly`) {
		t.Error("expected readonly attribute")
	}
}

func TestRequiredAttribute(t *testing.T) {
	html := render(t, Props{
		Name:     "bio",
		Required: true,
	})

	if !strings.Contains(html, ` required`) {
		t.Error("expected required attribute")
	}
}

func TestPlaceholder(t *testing.T) {
	html := render(t, Props{
		Name:        "bio",
		Placeholder: "Enter text here",
	})

	if !strings.Contains(html, `placeholder="Enter text here"`) {
		t.Error("expected placeholder attribute")
	}
}

func TestCustomID(t *testing.T) {
	html := render(t, Props{
		Name: "bio",
		ID:   "custom-id",
	})

	if !strings.Contains(html, `id="custom-id"`) {
		t.Error("expected custom id")
	}
	if !strings.Contains(html, `name="bio"`) {
		t.Error("expected name to remain unchanged")
	}
}

func TestCustomClass(t *testing.T) {
	html := render(t, Props{
		Name:  "bio",
		Class: "my-custom-class",
	})

	if !strings.Contains(html, `class="my-custom-class"`) {
		t.Error("expected custom class")
	}
}
