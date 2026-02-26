package command

import (
	"fmt"
	"os"
	"os/exec"
)

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
