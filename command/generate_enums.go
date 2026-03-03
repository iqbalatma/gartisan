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

	GenerateStatic(filepath.Join(enumBaseDir, "response_code.go"), templates.ResponseCodeTmpl)
}
