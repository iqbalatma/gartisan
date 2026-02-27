package command

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallMysql() {
	fmt.Println("Installing library gorm.io/driver/mysql")

	cmd := exec.Command("go", "get", "gorm.io/driver/mysql")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error installing mysql library")
	}
}
