package head

import "github.com/a-h/templ"

// Props contains configuration options for the Head component.
type Props struct {
	Title              string          // Page title
	Description        string          // Meta description
	IncludePico        bool            // Include Pico CSS (default true)
	IncludeHTMX        bool            // Include HTMX (default true)
	IncludeHyperscript bool            // Include _hyperscript (default true)
	ExtraHead          templ.Component // Additional head content
}

// DefaultProps returns Props with default values.
func DefaultProps() Props {
	return Props{
		IncludePico:        true,
		IncludeHTMX:        true,
		IncludeHyperscript: true,
	}
}

// picoCSS returns the embedded Pico CSS as a string.
func picoCSS() string {
	data, err := GetPicoCSS()
	if err != nil {
		return ""
	}
	return string(data)
}

// htmxJS returns the embedded HTMX JS as a string.
func htmxJS() string {
	data, err := GetHTMX()
	if err != nil {
		return ""
	}
	return string(data)
}

// hyperscriptJS returns the embedded _hyperscript JS as a string.
func hyperscriptJS() string {
	data, err := GetHyperscript()
	if err != nil {
		return ""
	}
	return string(data)
}

// rawStyle creates a raw style element with the given CSS content.
func rawStyle(css string) templ.Component {
	return templ.Raw("<style>" + css + "</style>")
}

// rawScript creates a raw script element with the given JS content.
func rawScript(js string) templ.Component {
	return templ.Raw("<script>" + js + "</script>")
}
