package main

import (
	"fmt"
	"os"

	"github.com/iqbalatma/gartisan/command"
	"github.com/iqbalatma/gartisan/utils"
)

var commands = map[string]func([]string){
	"make:controller": command.MakeController,
	"make:model":      command.MakeModel,
	"make:service":    command.MakeService,
	"make:repository": command.MakeRepository,
	"generate:utils":  command.GenerateUtils,
	"generate:enums":  command.GenerateEnums,
}

func main() {
	arguments := os.Args

	if len(os.Args) < 2 {
		utils.PrintIntroduction()
		return
	}
	cmdName := arguments[1]

	if cmd, exists := commands[cmdName]; exists {
		cmd(arguments)
	} else {
		fmt.Printf(utils.ANSI_RED+"Command %s not found\n", cmdName)
	}
}
