package utils

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/iqbalatma/gartisan/templates"
)

var controllerBaseDir = filepath.Join("app", "controller")

func MakeController(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(ANSI_RED + "Missing required argument. You need to add controller name as argument. e.g UserController ")
		return
	}

	argument := ExtractArgument(arguments[2], controllerBaseDir)
	MakeDirectoryIfNotExists(argument.FullPath)

	//map variable for template file
	data := map[string]string{
		"package_name":    "controller",
		"controller_name": argument.Name,
	}

	tmpl, err := template.New("controller").Parse(templates.ControllerTmpl)
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
		return
	}

	//write executed template file into buffer
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
		return
	}

	//write buffer into file
	err = SafeCreateFile(filepath.Join(argument.FullPath, argument.FileName), buffer.String())
	if err != nil {
		fmt.Printf(ANSI_RED+"Failed create controller %s \n", err.Error())
		return
	}

	fmt.Printf(ANSI_GREEN+"Successfully create controller %s \n", argument.Name)
	fmt.Printf(ANSI_RESET)
}
