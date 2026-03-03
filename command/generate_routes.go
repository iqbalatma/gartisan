package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
)

var routesBaseDir = "routes"

func GenerateRoutes(arguments []string) {
	fmt.Println("Generating routes")
	//create directory
	if err := os.MkdirAll(routesBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create routes directory")
		return
	}

	GenerateWithModuleName(filepath.Join(routesBaseDir, "route.go"), templates.RouteTmpl)
	fmt.Println("Generate routes successfully")
}
