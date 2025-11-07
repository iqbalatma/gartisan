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

	//set initiate data
	controllerNameArgument := arguments[2]
	controllerName := ToCamelCase(filepath.Base(controllerNameArgument))
	fileName := filepath.Base(controllerNameArgument) + ".go"
	pathExtension := filepath.Dir(controllerNameArgument)
	fullPath := controllerBaseDir
	if pathExtension != "." {
		fullPath = filepath.Join(fullPath, pathExtension)
	}
	fmt.Println(pathExtension)
	fmt.Println(fullPath)
	MakeDirectoryIfNotExists(fullPath)

	//get template file
	templateFile, err := template.ParseFiles(filepath.Join("templates", "controller.tmpl"))
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
		return
	}

	//map variable for template file
	data := map[string]string{
		"package_name":    "controller",
		"controller_name": controllerName,
	}

	//write executed template file into buffer
	var buffer bytes.Buffer
	err = templateFile.Execute(&buffer, data)
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
	}

	//write buffer into file
	err = SafeCreateFile(filepath.Join(fullPath, fileName), buffer.String())
	if err != nil {
		fmt.Printf(ANSI_RED+"Failed create controller %s \n", err.Error())
		return
	}

	fmt.Printf(ANSI_GREEN+"Successfully create controller %s \n", controllerName)
	fmt.Printf(ANSI_RESET)
}
