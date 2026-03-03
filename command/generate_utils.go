package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
)

var utilsBaseDir = "utils"
var modelsBaseDir = filepath.Join("app", "models")

func GenerateUtils(arguments []string) {
	fmt.Println("Generating utils")
	InstallBcrypt()
	InstallGin()
	InstallUUID()
	GenerateEnums(arguments)

	//create directory
	if err := os.MkdirAll(utilsBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create utils directory")
		return
	}

	if err := os.MkdirAll(modelsBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create utils directory")
		return
	}

	GenerateStatic(filepath.Join(utilsBaseDir, "hashing.go"), templates.HashingTmpl)
	GenerateStatic(filepath.Join(modelsBaseDir, "base_model.go"), templates.BaseModelTmpl)
	GenerateWithModuleName(filepath.Join(utilsBaseDir, "http_response.go"), templates.HttpResponseTmpl)
	GenerateWithModuleName(filepath.Join(utilsBaseDir, "paginate.go"), templates.PaginationTmpl)

	GoModTidy()
	fmt.Println("Generate utils successfully")
}
