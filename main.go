package main

import (
	"fmt"
	"os"
)

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {

	screen, err := NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	HandleEvents(screen)
}
