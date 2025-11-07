package utils

import (
	"fmt"
	"os"
)

func SafeCreateFile(path string, content string) error {
	// os.O_EXCL: memastikan operasi gagal jika file sudah ada
	// os.O_CREATE: membuat file jika tidak ada
	// 0644: izin standar file

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
