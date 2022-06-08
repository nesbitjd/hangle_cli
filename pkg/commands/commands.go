package commands

import (
	"fmt"
	"os"
)

// Handler handles the subcommand and flag inputs
func Handler() {

	if len(os.Args) < 2 {
		fmt.Println("expected 'play' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "play":
		play()
	default:
		fmt.Println("expected 'play' subcommands")
		os.Exit(1)
	}

}
