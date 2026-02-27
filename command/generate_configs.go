package command

import (
	"fmt"
	"os"

	"github.com/iqbalatma/gartisan/templates"
)

var configBaseDir = "config"

func GenerateConfigs(arguments []string) {
	fmt.Println("Generating config")
	InstallGodotenv()
	InstallGorm()
	InstallMysql()
	if err := os.MkdirAll(configBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create config directory")
		return
	}

	writeConfig()
	writeDatabase()
	fmt.Println("Generate config successfully")
}

func writeConfig() {
	fmt.Println("Writing config")
	err := os.WriteFile(configBaseDir+"/config.go", []byte(templates.ConfigTmpl), 0644)
	if err != nil {
		fmt.Println("Failed writing config.go:", err)
		return
	}
}

func writeDatabase() {
	fmt.Println("Writing database")
	err := os.WriteFile(configBaseDir+"/database.go", []byte(templates.DatabaseTmpl), 0644)
	if err != nil {
		fmt.Println("Failed writing database.go:", err)
		return
	}
}
