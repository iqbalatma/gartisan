package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
)

var utilsBaseDir = "utils"

func GenerateUtils(arguments []string) {
	fmt.Println("Generating utils")
	InstallBcrypt()
	InstallGin()
	GenerateEnums(arguments)

	//create directory
	if err := os.MkdirAll(utilsBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create utils directory")
		return
	}

	GenerateStatic(filepath.Join(utilsBaseDir, "hashing.go"), templates.HashingTmpl)
	GenerateWithModuleName(filepath.Join(utilsBaseDir, "http_response.go"), templates.HttpResponseTmpl)
	GenerateWithModuleName(filepath.Join(utilsBaseDir, "paginate.go"), templates.PaginationTmpl)

	GoModTidy()
	fmt.Println("Generate utils successfully")
}
