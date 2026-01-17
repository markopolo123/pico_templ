// Package button provides a Button component for Pico CSS.
package button

import "github.com/a-h/templ"

// Variant constants for button styling.
const (
	Primary   = ""          // Default primary style
	Secondary = "secondary" // Secondary style
	Contrast  = "contrast"  // Contrast style
)

// Props defines the properties for the Button component.
type Props struct {
	Text     string // Button text content
	Type     string // button, submit, reset (default: button)
	Variant  string // empty=primary, secondary, contrast
	Outline  bool   // Add .outline class
	Disabled bool   // Disabled state
	Class    string // Additional CSS classes
	// HTMX bindings
	HxGet     string
	HxPost    string
	HxPut     string
	HxDelete  string
	HxPatch   string
	HxTarget  string
	HxSwap    string
	HxTrigger string
	Attrs     templ.Attributes // Arbitrary additional attributes
}

// buttonType returns the button type, defaulting to "button".
func (p Props) buttonType() string {
	if p.Type == "" {
		return "button"
	}
	return p.Type
}

// classes builds the CSS class string for the button.
func (p Props) classes() string {
	var result string
	if p.Outline {
		result = "outline"
	}
	if p.Variant != "" {
		if result != "" {
			result += " " + p.Variant
		} else {
			result = p.Variant
		}
	}
	if p.Class != "" {
		if result != "" {
			result += " " + p.Class
		} else {
			result = p.Class
		}
	}
	return result
}
