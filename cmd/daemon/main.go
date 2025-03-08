package main

import (
	"fmt"
	"os"

	gitemoji "github.com/casper9429-kth/git-emoji-terminal"
)

func main() {
	// Show the emoji picker
	emoji := gitemoji.ShowPicker()

	// If an emoji was selected, print it
	if emoji != "" {
		fmt.Printf("Selected emoji: %s (copied to clipboard)\n", emoji)
		os.Exit(0)
	} else {
		fmt.Println("No emoji selected")
		os.Exit(1)
	}
}
