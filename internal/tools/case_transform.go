package tools

import (
	"strings"
	"unicode"
)

// CamelToPascal converts snake_case to PascalCase
func CamelToPascal(s string) string {
	// If the string is empty, return as it is
	if len(s) == 0 {
		return s
	}

	// Convert the first letter to uppercase and append the rest
	return strings.ToUpper(string(s[0])) + s[1:]
}

// IsCamelCase is a utility function to check if a string is camelCase
func IsCamelCase(s string) bool {
	// A camelCase string starts with a lowercase letter
	return len(s) > 0 && unicode.IsLower(rune(s[0]))
}

// PascalToSnake Converts PascalCase to snake_case
func PascalToSnake(s string) string {
	var result []rune

	for i, r := range s {
		// If it's an uppercase letter (and not the first character), prepend an underscore
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			// Append the lowercase version of the uppercase letter
			result = append(result, unicode.ToLower(r))
		} else {
			// Append the character as it is
			result = append(result, r)
		}
	}

	return string(result)
}

// IsPlural is a utility function to check if a string is plural
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

// PascalToCamel Convert PascalCase to camelCase, handling acronyms like 'UserID' to 'userId'
func PascalToCamel(input string) string {
	if len(input) == 0 {
		return input
	}

	// Convert the first rune to lowercase
	runes := []rune(input)

	// Find the first non-uppercase letter or the end of the string
	for i := 1; i < len(runes); i++ {
		// If the next letter is not uppercase, stop
		if !unicode.IsUpper(runes[i]) {
			runes[0] = unicode.ToLower(runes[0])
			break
		}
		// If we reach the last uppercase letter of an acronym (like 'ID'), convert the first letter and stop
		if i == len(runes)-1 {
			runes[0] = unicode.ToLower(runes[0])
		}
	}

	return string(runes)
}

// CamelToSnake Converts camelCase to snake_case
func CamelToSnake(input string) string {
	var result []rune

	for i, r := range input {
		// If it's an uppercase letter (and not the first character), prepend an underscore
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			// Append the lowercase version of the uppercase letter
			result = append(result, unicode.ToLower(r))
		} else {
			// Append the character as it is
			result = append(result, r)
		}
	}
	return string(result)
}
