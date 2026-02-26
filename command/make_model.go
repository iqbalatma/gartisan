package command

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
	"github.com/iqbalatma/gartisan/utils"
)

var modelBaseDir = filepath.Join("app", "model")

func MakeModel(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(utils.ANSI_RED + "Missing required argument. You need to add model name as argument. e.g User ")
		return
	}

	argument := utils.ExtractArgument(arguments[2], modelBaseDir)
	utils.MakeDirectoryIfNotExists(argument.FullPath)

	//map variable for template file
	data := map[string]string{
		"package_name": "model",
		"model_name":   argument.Name,
	}

	tmpl, err := template.New("model").Parse(templates.ModelTmpl)
	if err != nil {
		fmt.Println(utils.ANSI_RED + err.Error())
		return
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		fmt.Println(utils.ANSI_RED + err.Error())
		return
	}

	//write buffer into file
	err = utils.SafeCreateFile(filepath.Join(argument.FullPath, argument.FileName), buffer.String())
	if err != nil {
		fmt.Printf(utils.ANSI_RED+"Failed create model %s \n", err.Error())
		return
	}

	fmt.Printf(utils.ANSI_GREEN+"Successfully create model %s \n", argument.Name)
	fmt.Printf(utils.ANSI_RESET)
}
