package utils

import (
	"bytes"
	"fmt"
	"path/filepath"
	template2 "text/template"
)

var modelBaseDir = filepath.Join("app", "model")

func MakeModel(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(ANSI_RED + "Missing required argument. You need to add controller name as argument. e.g UserController ")
		return
	}

	argument := ExtractArgument(arguments[2], modelBaseDir)
	MakeDirectoryIfNotExists(argument.FullPath)

	templateFile, err := template2.ParseFiles(filepath.Join("templates", "model.tmpl"))
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
		return
	}

	//map variable for template file
	data := map[string]string{
		"package_name": "model",
		"model_name":   argument.Name,
	}

	var buffer bytes.Buffer
	err = templateFile.Execute(&buffer, data)
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
	}

	//write buffer into file
	err = SafeCreateFile(filepath.Join(argument.FullPath, argument.FileName), buffer.String())
	if err != nil {
		fmt.Printf(ANSI_RED+"Failed create model %s \n", err.Error())
		return
	}

	fmt.Printf(ANSI_GREEN+"Successfully create model %s \n", argument.Name)
	fmt.Printf(ANSI_RESET)
}
