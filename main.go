package main

import (
	"fmt"
	"os"

	"github.com/dullaz/freshdocs/cmd"
)

var version = "dev"

func main() {
	// Add version flag
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("FreshDocs version %s\n", version)
		os.Exit(0)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
