package command

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallLogrus() {
	fmt.Println("Installing library github.com/sirupsen/logrus")

	cmd := exec.Command("go", "get", "github.com/sirupsen/logrus")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error installing logrus library")
	}
}
