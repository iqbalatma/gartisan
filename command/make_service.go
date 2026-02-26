package command

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/iqbalatma/gartisan/templates"
	"github.com/iqbalatma/gartisan/utils"
)

var serviceBaseDir = filepath.Join("app", "service")

func MakeService(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(utils.ANSI_RED + "Missing required argument. You need to add service name as argument. e.g UserService ")
		return
	}

	argument := utils.ExtractArgument(arguments[2], serviceBaseDir)
	utils.MakeDirectoryIfNotExists(argument.FullPath)

	//map variable for template file
	data := map[string]string{
		"package_name": "service",
		"service_name": argument.Name,
	}

	tmpl, err := template.New("service").Parse(templates.ServiceTmpl)
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
		fmt.Printf(utils.ANSI_RED+"Failed create service %s \n", err.Error())
		return
	}

	fmt.Printf(utils.ANSI_GREEN+"Successfully create service %s \n", argument.Name)
	fmt.Printf(utils.ANSI_RESET)
}
