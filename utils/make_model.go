package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
)

var modelBaseDir = filepath.Join("app", "model")

func MakeModel(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(ANSI_RED + "Missing required argument. You need to add model name as argument. e.g User ")
		return
	}

	argument := ExtractArgument(arguments[2], modelBaseDir)
	MakeDirectoryIfNotExists(argument.FullPath)

	//map variable for template file
	data := map[string]string{
		"package_name": "model",
		"model_name":   argument.Name,
	}

	tmpl, err := template.New("model").Parse(templates.ModelTmpl)
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
		return
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
		return
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
