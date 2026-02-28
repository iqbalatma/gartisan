package command

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
	"github.com/iqbalatma/gartisan/utils"
)

var validatorBaseDir = filepath.Join("app", "validator")

func GenerateValidators(arguments []string) {
	fmt.Println("Generating validator")
	if err := os.MkdirAll(validatorBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create validator directory")
		return
	}
	moduleName, _ := utils.GetModuleName()
	templatePlaceholder := TemplatePlaceholder{
		ModuleName: moduleName,
	}
	writeValidator(templatePlaceholder)
	writeRegisterUniqueColumnValidator(templatePlaceholder)
	fmt.Println("Generate validator successfully")
}

func writeRegisterUniqueColumnValidator(templatePlaceholder TemplatePlaceholder) {
	fmt.Println("Writing validator")
	tmpl, err := template.New("validator").Parse(templates.RegisterUniqueColumnValidatorTmpl)
	if err != nil {
		fmt.Println("Failed writing register_unique_column_validator.go:", err)
		return
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templatePlaceholder)
	if err != nil {
		fmt.Println("Failed writing register_unique_column_validator.go:", err)
	}
	err = os.WriteFile(validatorBaseDir+"/register_unique_column_validator.go", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Failed writing register_unique_column_validator.go:", err)
		return
	}
}

func writeValidator(templatePlaceholder TemplatePlaceholder) {
	fmt.Println("Writing validator")
	tmpl, err := template.New("validator").Parse(templates.ValidatorTmpl)

	if err != nil {
		fmt.Println("Failed writing validator.go:", err)
		return
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templatePlaceholder)
	if err != nil {
		fmt.Println("Failed writing validator.go:", err)
	}
	err = os.WriteFile(validatorBaseDir+"/validator.go", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Failed writing validator.go:", err)
		return
	}
}
