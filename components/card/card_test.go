package card_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/markopolo123/pico_templ/components/card"
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

func TestCard_RendersArticle(t *testing.T) {
	html := render(t, card.Card(card.Props{}))
	if !strings.Contains(html, "<article") {
		t.Error("Card should render an <article> element")
	}
	if !strings.Contains(html, "</article>") {
		t.Error("Card should have closing </article> tag")
	}
}

func TestCardHeader_RendersHeader(t *testing.T) {
	html := render(t, card.CardHeader(card.HeaderProps{}))
	if !strings.Contains(html, "<header") {
		t.Error("CardHeader should render a <header> element")
	}
	if !strings.Contains(html, "</header>") {
		t.Error("CardHeader should have closing </header> tag")
	}
}

func TestCardFooter_RendersFooter(t *testing.T) {
	html := render(t, card.CardFooter(card.FooterProps{}))
	if !strings.Contains(html, "<footer") {
		t.Error("CardFooter should render a <footer> element")
	}
	if !strings.Contains(html, "</footer>") {
		t.Error("CardFooter should have closing </footer> tag")
	}
}

func TestCard_WithChildren(t *testing.T) {
	child := templ.Raw("<p>Test content</p>")
	component := card.Card(card.Props{})

	var buf bytes.Buffer
	err := component.Render(templ.WithChildren(context.Background(), child), &buf)
	if err != nil {
		t.Fatalf("failed to render component: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "<p>Test content</p>") {
		t.Error("Card should render children")
	}
}

func TestCardHeader_WithChildren(t *testing.T) {
	child := templ.Raw("<h3>Header Title</h3>")
	component := card.CardHeader(card.HeaderProps{})

	var buf bytes.Buffer
	err := component.Render(templ.WithChildren(context.Background(), child), &buf)
	if err != nil {
		t.Fatalf("failed to render component: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "<h3>Header Title</h3>") {
		t.Error("CardHeader should render children")
	}
}

func TestCardFooter_WithChildren(t *testing.T) {
	child := templ.Raw("<button>Action</button>")
	component := card.CardFooter(card.FooterProps{})

	var buf bytes.Buffer
	err := component.Render(templ.WithChildren(context.Background(), child), &buf)
	if err != nil {
		t.Fatalf("failed to render component: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "<button>Action</button>") {
		t.Error("CardFooter should render children")
	}
}

func TestCard_CustomClass(t *testing.T) {
	html := render(t, card.Card(card.Props{Class: "custom-class"}))
	if !strings.Contains(html, `class="custom-class"`) {
		t.Error("Card should apply custom class")
	}
}

func TestCardHeader_CustomClass(t *testing.T) {
	html := render(t, card.CardHeader(card.HeaderProps{Class: "header-class"}))
	if !strings.Contains(html, `class="header-class"`) {
		t.Error("CardHeader should apply custom class")
	}
}

func TestCardFooter_CustomClass(t *testing.T) {
	html := render(t, card.CardFooter(card.FooterProps{Class: "footer-class"}))
	if !strings.Contains(html, `class="footer-class"`) {
		t.Error("CardFooter should apply custom class")
	}
}

func TestCard_AttrsSpread(t *testing.T) {
	attrs := templ.Attributes{
		"id":          "test-card",
		"data-testid": "card-component",
	}
	html := render(t, card.Card(card.Props{Attrs: attrs}))

	if !strings.Contains(html, `id="test-card"`) {
		t.Error("Card should spread id attribute")
	}
	if !strings.Contains(html, `data-testid="card-component"`) {
		t.Error("Card should spread data-testid attribute")
	}
}
