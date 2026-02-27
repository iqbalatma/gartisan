package command

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallGorm() {
	fmt.Println("Installing library gorm.io/gorm")

	cmd := exec.Command("go", "get", "gorm.io/gorm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error installing gorm library")
	}
}
