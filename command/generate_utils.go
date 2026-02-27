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
	InstallGin()
	GenerateEnums(arguments)
	moduleName, _ := utils.GetModuleName()
	templatePlaceholder := TemplatePlaceholder{
		ModuleName: moduleName,
	}

	//create directory
	if err := os.MkdirAll(utilsBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create utils directory")
		return
	}

	//write hashing
	writeHashing()
	writeHttpResponse(templatePlaceholder)
	writePaginate(templatePlaceholder)

	GoModTidy()
	fmt.Println("Generate utils successfully")
}

func writeHashing() {
	fmt.Println("Writing hashing")

	err := os.WriteFile(utilsBaseDir+"/hashing.go", []byte(templates.HashingTmpl), 0644)
	if err != nil {
		fmt.Println("Failed writing hashing.go:", err)
		return
	}
}

func writeHttpResponse(templatePlaceholder TemplatePlaceholder) {
	fmt.Println("Writing http response")

	tmpl, err := template.New("http").Parse(templates.HttpResponseTmpl)
	if err != nil {
		fmt.Println("Failed writing http_response.go:", err)
		return
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templatePlaceholder)
	if err != nil {
		fmt.Println("Failed writing http_response.go:", err)
	}
	err = os.WriteFile(utilsBaseDir+"/http_response.go", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Failed writing http_response.go:", err)
		return
	}
}

func writePaginate(templatePlaceholder TemplatePlaceholder) {
	fmt.Println("Writing paginate")
	tmpl, err := template.New("paginate").Parse(templates.PaginationTmpl)
	if err != nil {
		fmt.Println("Failed writing paginate.go:", err)
		return
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templatePlaceholder)
	if err != nil {
		fmt.Println("Failed writing http_response.go:", err)
	}
	err = os.WriteFile(utilsBaseDir+"/paginate.go", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Failed writing paginate.go:", err)
		return
	}
}
