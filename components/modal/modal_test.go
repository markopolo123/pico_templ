package modal

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

func TestModal_RendersDialogWithCorrectID(t *testing.T) {
	html := render(t, Modal(Props{ID: "test-modal"}))

	if !strings.Contains(html, `<dialog`) {
		t.Error("expected <dialog> element")
	}
	if !strings.Contains(html, `id="test-modal"`) {
		t.Error("expected id=\"test-modal\" attribute")
	}
}

func TestModal_OpenAttributeRendersWhenTrue(t *testing.T) {
	html := render(t, Modal(Props{ID: "open-modal", Open: true}))

	if !strings.Contains(html, ` open`) && !strings.Contains(html, ` open `) && !strings.Contains(html, ` open>`) {
		t.Error("expected open attribute when Open=true")
	}
}

func TestModal_OpenAttributeAbsentWhenFalse(t *testing.T) {
	html := render(t, Modal(Props{ID: "closed-modal", Open: false}))

	// Check that 'open' doesn't appear as an attribute
	if strings.Contains(html, ` open `) || strings.Contains(html, ` open>`) {
		t.Error("expected no open attribute when Open=false")
	}
}

func TestModal_HasArticleWrapper(t *testing.T) {
	html := render(t, Modal(Props{ID: "article-modal"}))

	if !strings.Contains(html, "<article>") {
		t.Error("expected <article> element inside dialog")
	}
}

func TestModal_HasClickOutsideHyperscript(t *testing.T) {
	html := render(t, Modal(Props{ID: "click-modal"}))

	if !strings.Contains(html, `_="on click if event.target === me call me.close()"`) {
		t.Error("expected _hyperscript for click-outside-to-close")
	}
}

func TestModal_AdditionalClasses(t *testing.T) {
	html := render(t, Modal(Props{ID: "class-modal", Class: "custom-class"}))

	if !strings.Contains(html, `class="custom-class"`) {
		t.Error("expected class attribute with custom-class")
	}
}

func TestModalHeader_RendersTitle(t *testing.T) {
	html := render(t, ModalHeader(HeaderProps{Title: "Test Title"}))

	if !strings.Contains(html, "<header>") {
		t.Error("expected <header> element")
	}
	if !strings.Contains(html, "Test Title") {
		t.Error("expected title text")
	}
	if !strings.Contains(html, "<strong>") {
		t.Error("expected title wrapped in <strong>")
	}
}

func TestModalHeader_RendersCloseButton(t *testing.T) {
	html := render(t, ModalHeader(HeaderProps{Title: "Test", ShowClose: true}))

	if !strings.Contains(html, `aria-label="Close"`) {
		t.Error("expected close button with aria-label")
	}
	if !strings.Contains(html, `rel="prev"`) {
		t.Error("expected close button with rel=\"prev\"")
	}
	// HTML escapes < and > in attribute values
	if !strings.Contains(html, `_="on click call closest &lt;dialog/&gt;.close()"`) {
		t.Error("expected close button _hyperscript")
	}
}

func TestModalHeader_ShowCloseDefaultsWithTitle(t *testing.T) {
	html := render(t, ModalHeader(HeaderProps{Title: "Has Title"}))

	if !strings.Contains(html, `aria-label="Close"`) {
		t.Error("expected close button when title is set")
	}
}

func TestModalFooter_RendersFooter(t *testing.T) {
	html := render(t, ModalFooter(FooterProps{}))

	if !strings.Contains(html, "<footer>") {
		t.Error("expected <footer> element")
	}
}

func TestModalFooter_AdditionalClasses(t *testing.T) {
	html := render(t, ModalFooter(FooterProps{Class: "footer-class"}))

	if !strings.Contains(html, `class="footer-class"`) {
		t.Error("expected class attribute")
	}
}

func TestModalTrigger_HasCorrectHyperscript(t *testing.T) {
	html := render(t, ModalTrigger(TriggerProps{ModalID: "my-modal", Text: "Open"}))

	if !strings.Contains(html, `_="on click call #my-modal.showModal()"`) {
		t.Error("expected _hyperscript to call showModal on modal ID")
	}
}

func TestModalTrigger_RendersButtonText(t *testing.T) {
	html := render(t, ModalTrigger(TriggerProps{ModalID: "modal", Text: "Open Modal"}))

	if !strings.Contains(html, "<button") {
		t.Error("expected <button> element")
	}
	if !strings.Contains(html, "Open Modal") {
		t.Error("expected button text")
	}
}

func TestModalTrigger_Variant(t *testing.T) {
	html := render(t, ModalTrigger(TriggerProps{ModalID: "modal", Text: "Test", Variant: "secondary"}))

	if !strings.Contains(html, `class="secondary"`) {
		t.Error("expected secondary class")
	}
}

func TestModalClose_HasCorrectHyperscriptWithModalID(t *testing.T) {
	html := render(t, ModalClose(CloseProps{ModalID: "close-modal", Text: "Close"}))

	if !strings.Contains(html, `_="on click call #close-modal.close()"`) {
		t.Error("expected _hyperscript to call close on modal ID")
	}
}

func TestModalClose_HasCorrectHyperscriptWithoutModalID(t *testing.T) {
	html := render(t, ModalClose(CloseProps{Text: "Close"}))

	// HTML escapes < and > in attribute values
	if !strings.Contains(html, `_="on click call closest &lt;dialog/&gt;.close()"`) {
		t.Error("expected _hyperscript to call close on closest dialog")
	}
}

func TestModalClose_RendersButtonText(t *testing.T) {
	html := render(t, ModalClose(CloseProps{Text: "Cancel"}))

	if !strings.Contains(html, "<button") {
		t.Error("expected <button> element")
	}
	if !strings.Contains(html, "Cancel") {
		t.Error("expected button text")
	}
}

func TestModalClose_Variant(t *testing.T) {
	html := render(t, ModalClose(CloseProps{Text: "Cancel", Variant: "secondary"}))

	if !strings.Contains(html, `class="secondary"`) {
		t.Error("expected secondary class")
	}
}
