package utils

import (
	"path/filepath"
	"strings"
	"unicode"
)

type ArgumentInfo struct {
	Name          string
	FileName      string
	PathExtension string
	FullPath      string
}

func ToCamelCase(source string) string {
	explodedString := strings.Split(source, "_")

	result := ""
	for _, word := range explodedString {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])

		result += string(runes)
	}

	return result
}

func ExtractArgument(argument string) *ArgumentInfo {
	pathExtension := filepath.Dir(argument)
	fullPath := ""
	if pathExtension != "." {
		fullPath = filepath.Join(fullPath, pathExtension)
	}
	return &ArgumentInfo{
		Name:          ToCamelCase(filepath.Base(argument)),
		FileName:      filepath.Base(argument) + ".go",
		PathExtension: pathExtension,
		FullPath:      fullPath,
	}
}
