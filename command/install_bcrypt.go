package command

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallBcrypt() {
	fmt.Println("Installing library golang.org/x/crypto/bcrypt")

	cmd := exec.Command("go", "get", "golang.org/x/crypto/bcrypt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error installing golang bcrypt library")
	}
}
