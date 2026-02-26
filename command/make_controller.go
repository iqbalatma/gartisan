package command

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/iqbalatma/gartisan/templates"
	"github.com/iqbalatma/gartisan/utils"
)

var controllerBaseDir = filepath.Join("app", "controller")

func MakeController(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(utils.ANSI_RED + "Missing required argument. You need to add controller name as argument. e.g UserController ")
		return
	}

	argument := utils.ExtractArgument(arguments[2], controllerBaseDir)
	utils.MakeDirectoryIfNotExists(argument.FullPath)

	//map variable for template file
	data := map[string]string{
		"package_name":    "controller",
		"controller_name": argument.Name,
	}

	tmpl, err := template.New("controller").Parse(templates.ControllerTmpl)
	if err != nil {
		fmt.Println(utils.ANSI_RED + err.Error())
		return
	}

	//write executed template file into buffer
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		fmt.Println(utils.ANSI_RED + err.Error())
		return
	}

	//write buffer into file
	err = utils.SafeCreateFile(filepath.Join(argument.FullPath, argument.FileName), buffer.String())
	if err != nil {
		fmt.Printf(utils.ANSI_RED+"Failed create controller %s \n", err.Error())
		return
	}

	fmt.Printf(utils.ANSI_GREEN+"Successfully create controller %s \n", argument.Name)
	fmt.Printf(utils.ANSI_RESET)
}
