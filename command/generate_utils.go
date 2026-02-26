package command

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/iqbalatma/gartisan/templates"
	"github.com/iqbalatma/gartisan/utils"
)

var utilsBaseDir = "utils"

func GenerateUtils(arguments []string) {
	fmt.Println("Generating utils")
	InstallBcrypt()

	if err := os.MkdirAll(utilsBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create utils directory")
		return
	}

	err := os.WriteFile(utilsBaseDir+"/hashing.go", []byte(templates.HashingTmpl), 0644)
	if err != nil {
		fmt.Println("Failed writing hashing.go:", err)
		return
	}

	tmpl, err := template.New("http").Parse(templates.HttpResponseTmpl)
	if err != nil {
		fmt.Println("Failed writing http_response.go:", err)
		return
	}

	moduleName, _ := utils.GetModuleName()
	data := struct {
		ModuleName string
	}{
		ModuleName: moduleName,
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		fmt.Println("Failed writing http_response.go:", err)
	}
	err = os.WriteFile(utilsBaseDir+"/http_response.go", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Failed writing http_response.go:", err)
		return
	}
	
	GoModTidy()
	fmt.Println("Generate utils successfully")
}
