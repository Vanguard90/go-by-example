// Define a package name for this package
package main

// Import fmt package, which contains functions for formatting text,
// including printing to the console (standard library)
import (
	"fmt"

	"rsc.io/quote"
)

// A main function executes by default when you run the main package
func main() {
    fmt.Println(quote.Go())
}

// Run with: go run .