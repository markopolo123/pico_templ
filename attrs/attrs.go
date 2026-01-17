// Package attrs provides common attribute helpers for pico_templ components.
package attrs

// HtmxAttrs contains HTMX attribute bindings for components.
type HtmxAttrs struct {
	Get      string // hx-get URL
	Post     string // hx-post URL
	Put      string // hx-put URL
	Delete   string // hx-delete URL
	Patch    string // hx-patch URL
	Target   string // hx-target selector
	Swap     string // hx-swap strategy
	Trigger  string // hx-trigger event
	Confirm  string // hx-confirm message
	Indicator string // hx-indicator selector
	PushURL  string // hx-push-url
	Select   string // hx-select selector
	Vals     string // hx-vals JSON
}

// HasHtmx returns true if any HTMX attributes are set.
func (h HtmxAttrs) HasHtmx() bool {
	return h.Get != "" || h.Post != "" || h.Put != "" || h.Delete != "" || h.Patch != ""
}
