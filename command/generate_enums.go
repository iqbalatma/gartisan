package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
)

var enumBaseDir = filepath.Join("app", "enums")

func GenerateEnums(arguments []string) {
	fmt.Println("Generating enums")

	if err := os.MkdirAll(enumBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create enum directory")
		return
	}

	err := os.WriteFile(enumBaseDir+"/response_code.go", []byte(templates.ResponseCodeTmpl), 0644)
	if err != nil {
		fmt.Println("Failed writing response_code.go:", err)
		return
	}
	fmt.Println("Generate enums successfully")
}
