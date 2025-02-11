package utils

import (
	"strings"
	"unicode"
)

func CamelToPascal(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}

func IsCamelCase(s string) bool {
	// A camelCase string starts with a lowercase letter
	return len(s) > 0 && unicode.IsLower(rune(s[0]))
}

func PascalToSnake(s string) string {
	var result []rune

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}

func IsPlural(word string) bool {
	word = strings.ToLower(word)
	if len(word) == 0 {
		return false
	}

	if strings.HasSuffix(word, "s") && !strings.HasSuffix(word, "ss") {
		return true
	}

	return false
}

func PascalToCamel(input string) string {
	if len(input) == 0 {
		return input
	}

	runes := []rune(input)

	for i := 1; i < len(runes); i++ {
		if !unicode.IsUpper(runes[i]) {
			runes[0] = unicode.ToLower(runes[0])
			break
		}
		if i == len(runes)-1 {
			runes[0] = unicode.ToLower(runes[0])
		}
	}

	return string(runes)
}

func CamelToSnake(input string) string {
	var result []rune

	for i, r := range input {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
