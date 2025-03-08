package main

import (
	"fmt"
	"os"

	"github.com/casper/gitemoji"
)

func main() {
	// Show the emoji picker
	emoji := gitemoji.ShowPicker()

	// If an emoji was selected, copy it to clipboard and print it
	if emoji != "" {
		fmt.Printf("Selected emoji: %s (copied to clipboard)\n", emoji)
		os.Exit(0)
	} else {
		fmt.Println("No emoji selected")
		os.Exit(1)
	}
}
