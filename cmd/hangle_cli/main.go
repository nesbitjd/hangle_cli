package main

import (
	"fmt"

	"github.com/nesbitjd/hangle_cli/pkg/commands"
)

func main() {

	fmt.Println("Let's Hangle!")

	err := commands.Handler()
	if err != nil {
		panic(err)
	}

}
