package command

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/iqbalatma/gartisan/templates"
	"github.com/iqbalatma/gartisan/utils"
)

var repositoryBaseDir = filepath.Join("app", "repository")

func MakeRepository(arguments []string) {
	if len(arguments) < 3 {
		fmt.Println(utils.ANSI_RED + "Missing required argument. You need to add repository name as argument. e.g UserRepository ")
		return
	}

	argument := utils.ExtractArgument(arguments[2], repositoryBaseDir)
	utils.MakeDirectoryIfNotExists(argument.FullPath)

	//map variable for template file
	data := map[string]string{
		"package_name":    "repository",
		"repository_name": argument.Name,
	}

	//get template file
	tmpl, err := template.New("repository").Parse(templates.RepositoryTmpl)
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
		fmt.Printf(utils.ANSI_RED+"Failed create repository %s \n", err.Error())
		return
	}

	fmt.Printf(utils.ANSI_GREEN+"Successfully create repository %s \n", argument.Name)
	fmt.Printf(utils.ANSI_RESET)
}
