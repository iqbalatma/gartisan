package command

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/iqbalatma/gartisan/utils"
)

func GenerateStatic(fullFileName string, templateContent string) {
	fmt.Printf("Writing %s \n", fullFileName)
	err := os.WriteFile(fullFileName, []byte(templateContent), 0644)
	if err != nil {
		fmt.Printf("Failed writing %s:%s", fullFileName, err)
		return
	}
}

func GenerateWithModuleName(fullFileName string, templateContent string) {
	moduleName, _ := utils.GetModuleName()
	templatePlaceholder := TemplatePlaceholder{
		ModuleName: moduleName,
	}

	fmt.Printf("Writing validator %s \n", fullFileName)
	tmpl, err := template.New(fullFileName).Parse(templateContent)
	if err != nil {
		fmt.Printf("Failed writing %s:%s", fullFileName, err)
		return
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templatePlaceholder)
	if err != nil {
		fmt.Printf("Failed writing %s:%s", fullFileName, err)
	}
	err = os.WriteFile(fullFileName, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Failed writing %s:%s", fullFileName, err)
		return
	}
}
