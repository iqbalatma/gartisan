package command

import (
	"fmt"
	"os"
	"os/exec"
)

func Install(library string) {
	fmt.Printf("Installing library %s\n", library)

	cmd := exec.Command("go", "get", library)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error installing library %s \n", library)
	}
}

func InstallGodotenv() {
	Install("github.com/joho/godotenv")
}

func InstallGorm() {
	Install("gorm.io/gorm")
}

func InstallLogrus() {
	Install("github.com/sirupsen/logrus")
}

func InstallMysql() {
	Install("gorm.io/driver/mysql")
}

func InstallGin() {
	Install("github.com/gin-gonic/gin")
}

func InstallBcrypt() {
	Install("golang.org/x/crypto/bcrypt")
}

func GoModTidy() {
	fmt.Println("Running go mod tidy")

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return
	}
}
