package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/markopolo123/pico_templ/docs/pages"
)

// Page represents a page to be generated
type Page struct {
	Name      string
	Filename  string
	Component templ.Component
}

func main() {
	distDir := "dist"

	// Create dist directory
	if err := os.MkdirAll(distDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating dist directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created output directory: %s\n", distDir)

	// Define pages to generate
	pagesToGenerate := []Page{
		{Name: "Index", Filename: "index.html", Component: pages.Index()},
		{Name: "Components", Filename: "components.html", Component: pages.Components()},
		{Name: "Forms", Filename: "forms.html", Component: pages.Forms()},
		{Name: "Layout", Filename: "layout.html", Component: pages.Layout()},
		{Name: "Content", Filename: "content.html", Component: pages.Content()},
	}

	// Generate each page
	var errors []error
	for _, page := range pagesToGenerate {
		outputPath := filepath.Join(distDir, page.Filename)
		fmt.Printf("Generating %s...", page.Name)

		if err := generatePage(outputPath, page.Component); err != nil {
			fmt.Printf(" FAILED: %v\n", err)
			errors = append(errors, fmt.Errorf("%s: %w", page.Name, err))
		} else {
			fmt.Printf(" OK (%s)\n", outputPath)
		}
	}

	// Report results
	fmt.Println()
	if len(errors) > 0 {
		fmt.Printf("Generation completed with %d error(s):\n", len(errors))
		for _, err := range errors {
			fmt.Fprintf(os.Stderr, "  - %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %d pages to %s/\n", len(pagesToGenerate), distDir)
}

// generatePage renders a templ component to an HTML file
func generatePage(path string, component templ.Component) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	if err := component.Render(context.Background(), f); err != nil {
		return fmt.Errorf("failed to render component: %w", err)
	}

	return nil
}
