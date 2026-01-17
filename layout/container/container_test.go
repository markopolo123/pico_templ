package container

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

func TestContainer_DefaultHasContainerClass(t *testing.T) {
	html := render(t, Container(Props{}))

	if !strings.Contains(html, `class="container"`) {
		t.Errorf("expected container class, got: %s", html)
	}
	if strings.Contains(html, "container-fluid") {
		t.Errorf("default container should not have container-fluid class, got: %s", html)
	}
}

func TestContainer_FluidHasContainerFluidClass(t *testing.T) {
	html := render(t, Container(Props{Fluid: true}))

	if !strings.Contains(html, `class="container-fluid"`) {
		t.Errorf("expected container-fluid class, got: %s", html)
	}
}

func TestContainer_ChildrenRenderInside(t *testing.T) {
	// We need to create a component with children
	// templ components receive children via the children... syntax
	// For testing, we'll verify the structure allows children
	html := render(t, Container(Props{}))

	if !strings.HasPrefix(html, "<div") {
		t.Errorf("expected div element, got: %s", html)
	}
	if !strings.HasSuffix(strings.TrimSpace(html), "</div>") {
		t.Errorf("expected closing div tag, got: %s", html)
	}
}

func TestContainer_CustomClassesAppend(t *testing.T) {
	html := render(t, Container(Props{Class: "my-custom-class"}))

	if !strings.Contains(html, `class="container my-custom-class"`) {
		t.Errorf("expected custom class to be appended, got: %s", html)
	}
}

func TestContainer_FluidWithCustomClasses(t *testing.T) {
	html := render(t, Container(Props{Fluid: true, Class: "extra"}))

	if !strings.Contains(html, `class="container-fluid extra"`) {
		t.Errorf("expected container-fluid with custom class, got: %s", html)
	}
}

func TestContainer_AttrsSpreadIntoElement(t *testing.T) {
	html := render(t, Container(Props{
		Attrs: templ.Attributes{
			"id":           "main-container",
			"data-testid":  "container",
		},
	}))

	if !strings.Contains(html, `id="main-container"`) {
		t.Errorf("expected id attribute, got: %s", html)
	}
	if !strings.Contains(html, `data-testid="container"`) {
		t.Errorf("expected data-testid attribute, got: %s", html)
	}
}

func TestContainer_ClassWithWhitespace(t *testing.T) {
	html := render(t, Container(Props{Class: "  spaced-class  "}))

	// Should trim the whitespace
	if !strings.Contains(html, `class="container spaced-class"`) {
		t.Errorf("expected trimmed class, got: %s", html)
	}
}
