# pico_templ

A Go [templ](https://templ.guide) component library for [Pico CSS](https://picocss.com).

## Features

- Full Pico CSS v2 component coverage
- Type-safe component props
- HTMX attribute bindings on interactive components
- [_hyperscript](https://hyperscript.org) for client-side interactivity (modals, accordions, dropdowns)
- Comprehensive unit tests

## Installation

```bash
go get github.com/markopolo123/pico_templ
```

## Requirements

- Go 1.23+
- [templ](https://templ.guide) CLI for code generation
- [just](https://github.com/casey/just) for task running (optional)

## Usage

```go
package main

import (
    "os"
    "context"

    "github.com/markopolo123/pico_templ/components/button"
)

func main() {
    btn := button.Button(button.Props{
        Text:    "Click me",
        Variant: button.Secondary,
        HxPost:  "/api/action",
        HxSwap:  "innerHTML",
    })
    btn.Render(context.Background(), os.Stdout)
}
```

## Components

### Components
- Button (with HTMX bindings)
- Modal (with _hyperscript)
- Accordion (with _hyperscript)
- Card
- Dropdown (with _hyperscript)
- Nav
- Progress

### Forms
- Input
- Textarea
- Select
- Checkbox
- Radio
- Switch
- Range

### Content
- Typography
- Table
- Link
- Loading

### Layout
- Container
- Grid
- Group
- Tooltip

## Development

```bash
# Generate templ files
just generate

# Run tests
just test

# Format code
just fmt

# Watch for changes
just watch
```

## License

MIT
