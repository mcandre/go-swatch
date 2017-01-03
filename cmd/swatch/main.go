// Package main provides a swatch executable.
package main

import (
	"fmt"

	"github.com/mcandre/go-swatch"
)

// main is the entrypoint for this application.
func main() {
	fmt.Println(swatch.Swatch())
}
