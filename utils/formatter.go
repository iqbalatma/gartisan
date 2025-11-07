package utils

import (
	"strings"
	"unicode"
)

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
