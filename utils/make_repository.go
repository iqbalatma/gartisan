package utils

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"
)

var repositoryBaseDir = filepath.Join("app", "repository")

func MakeRepository(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(ANSI_RED + "Missing required argument. You need to add repository name as argument. e.g UserRepository ")
		return
	}

	argument := ExtractArgument(arguments[2], repositoryBaseDir)
	MakeDirectoryIfNotExists(argument.FullPath)

	//get template file
	templateFile, err := template.ParseFiles(filepath.Join("templates", "repository.tmpl"))
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
		return
	}

	//map variable for template file
	data := map[string]string{
		"package_name":    "repository",
		"repository_name": argument.Name,
	}

	var buffer bytes.Buffer
	err = templateFile.Execute(&buffer, data)
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
	}

	//write buffer into file
	err = SafeCreateFile(filepath.Join(argument.FullPath, argument.FileName), buffer.String())
	if err != nil {
		fmt.Printf(ANSI_RED+"Failed create repository %s \n", err.Error())
		return
	}

	fmt.Printf(ANSI_GREEN+"Successfully create repository %s \n", argument.Name)
	fmt.Printf(ANSI_RESET)
}
