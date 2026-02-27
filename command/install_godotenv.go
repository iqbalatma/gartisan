package command

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallGodotenv() {
	fmt.Println("Installing library github.com/joho/godotenv")

	cmd := exec.Command("go", "get", "github.com/joho/godotenv")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error installing godotenv library")
	}
}
