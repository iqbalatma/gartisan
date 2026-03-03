package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
)

var errorsBaseDir = "errors"

func GenerateErrors(arguments []string) {
	fmt.Println("Generating errors")
	//create directory
	if err := os.MkdirAll(errorsBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create errors directory")
		return
	}

	GenerateWithModuleName(filepath.Join(errorsBaseDir, "exception.go"), templates.ExceptionTmpl)
	fmt.Println("Generate errors successfully")
}
