package command

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallGin() {
	fmt.Println("Installing library github.com/gin-gonic/gin")

	cmd := exec.Command("go", "get", "github.com/gin-gonic/gin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error installing golang gin library")
	}
}
