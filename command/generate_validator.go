package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
)

var validatorBaseDir = filepath.Join("app", "validator")

func GenerateValidators(arguments []string) {
	fmt.Println("Generating validator")
	if err := os.MkdirAll(validatorBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create validator directory")
		return
	}

	GenerateWithModuleName(filepath.Join(validatorBaseDir, "register_unique_column_validator.go"), templates.RegisterUniqueColumnValidatorTmpl)
	GenerateWithModuleName(filepath.Join(validatorBaseDir, "validator.go"), templates.ValidatorTmpl)
	fmt.Println("Generate validator successfully")
}
