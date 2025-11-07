package main

import (
	"fmt"
	"os"

	"github.com/iqbalatma/gartisan/utils"
)

func main() {
	arguments := os.Args

	if len(os.Args) < 2 {
		utils.PrintIntroduction()
		return
	}
	command := arguments[1]

	switch command {
	case "make:controller":
		utils.MakeController(arguments)
	default:
		fmt.Printf(utils.ANSI_RED+"Command %s not found\n", command)
	}
}
