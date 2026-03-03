package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iqbalatma/gartisan/templates"
)

var configBaseDir = "config"

func GenerateConfigs(arguments []string) {
	fmt.Println("Generating config")
	InstallGodotenv()
	InstallGorm()
	InstallMysql()
	InstallLogrus()
	if err := os.MkdirAll(configBaseDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create config directory")
		return
	}

	GenerateStatic(filepath.Join(configBaseDir, "config.go"), templates.ConfigTmpl)
	GenerateStatic(filepath.Join(configBaseDir, "database.go"), templates.DatabaseTmpl)
	GenerateStatic(filepath.Join(configBaseDir, "logger.go"), templates.LoggerTmpl)
	fmt.Println("Generate config successfully")
}
