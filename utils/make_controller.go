package utils

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"
)

var controllerBaseDir = filepath.Join("app", "controller")

func MakeController(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(ANSI_RED + "Missing required argument. You need to add controller name as argument. e.g UserController ")
		return
	}

	argument := ExtractArgument(arguments[2])
	MakeDirectoryIfNotExists(argument.FullPath)

	//get template file
	templateFile, err := template.ParseFiles(filepath.Join("templates", "controller.tmpl"))
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
		return
	}

	//map variable for template file
	data := map[string]string{
		"package_name":    "controller",
		"controller_name": argument.Name,
	}

	//write executed template file into buffer
	var buffer bytes.Buffer
	err = templateFile.Execute(&buffer, data)
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
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
