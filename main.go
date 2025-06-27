package main

import (
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	// Check if there are command line arguments
	if len(os.Args) < 2 {
		// No URL provided, silently exit
		return
	}

	// Combine all arguments into a single string
	// This handles URLs with spaces or special characters
	url := strings.Join(os.Args[1:], " ")

	// Write the URL to the clipboard
	err := clipboard.WriteAll(url)
	if err != nil {
		// Fail silently if clipboard write fails
		return
	}

	// Exit immediately after copying
}
