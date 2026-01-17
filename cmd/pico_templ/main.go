// Package main provides the pico_templ CLI for development utilities.
package main

import (
	"fmt"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("pico_templ %s (commit: %s, built: %s)\n", version, commit, date)
		return
	}

	fmt.Println("pico_templ - Pico CSS component library for Go templ")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("\nThis is a library package. Import it in your Go project:")
	fmt.Println("  import \"github.com/markopolo123/pico_templ/components/button\"")
	fmt.Println("\nFor documentation, visit: https://github.com/markopolo123/pico_templ")
}
