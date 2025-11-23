package utils

import (
	"fmt"
	"os"
)

func SafeCreateFile(path string, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func MakeDirectoryIfNotExists(fullPath string) {
	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		fmt.Println(ANSI_RED + err.Error())
	}
}
