package main

import (
	"fmt"
	"os"

	"github.com/iqbalatma/gartisan/command"
	"github.com/iqbalatma/gartisan/utils"
)

var commands = map[string]func([]string){
	"make:controller":     command.MakeController,
	"make:model":          command.MakeModel,
	"make:service":        command.MakeService,
	"make:repository":     command.MakeRepository,
	"make:migration":      command.MakeMigration,
	"generate:utils":      command.GenerateUtils,
	"generate:enums":      command.GenerateEnums,
	"generate:configs":    command.GenerateConfigs,
	"generate:validators": command.GenerateValidators,
	"generate:errors":     command.GenerateErrors,
	"generate:routes":     command.GenerateRoutes,
	"generate":            command.GenerateAll,
	"migrate":             command.MigrateUp,
	"migrate:up":          command.MigrateUp,
	"migrate:down":        command.MigrateDown,
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
