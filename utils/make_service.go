package utils

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/iqbalatma/gartisan/templates"
)

var serviceBaseDir = filepath.Join("app", "service")

func MakeService(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(ANSI_RED + "Missing required argument. You need to add service name as argument. e.g UserService ")
		return
	}

	argument := ExtractArgument(arguments[2], serviceBaseDir)
	MakeDirectoryIfNotExists(argument.FullPath)

	//map variable for template file
	data := map[string]string{
		"package_name": "service",
		"service_name": argument.Name,
	}

	tmpl, err := template.New("service").Parse(templates.ServiceTmpl)
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
		fmt.Printf(ANSI_RED+"Failed create service %s \n", err.Error())
		return
	}

	fmt.Printf(ANSI_GREEN+"Successfully create service %s \n", argument.Name)
	fmt.Printf(ANSI_RESET)
}
