package main

import (
	"flag"
	hg "hangman/functions"
)

func main() {

	hardMode := flag.Bool("hard", false, "Enable hard mode")

	flag.Parse() // Scan the options of the terminal

	if *hardMode { // Start the game with the mode

		hg.PlayHangman(true) // Call hard mode (go run main.go --hard)
	} else {
		hg.PlayHangman(false) // Call normal mode (go run main.go)
	}
}
