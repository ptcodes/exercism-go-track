package atbash

import (
	"regexp"
	"strings"
	"unicode"
)

func Atbash(input string) string {
	input = regexp.MustCompile(`[^\w\d]`).ReplaceAllString(input, "")
	input = strings.ToLower(input)

	encrypted := []rune{}

	for i, char := range input {
		if unicode.IsLetter(char) {
			char = 'a' + ('z' - char)
		}
		if i%5 == 0 {
			encrypted = append(encrypted, ' ')
		}
		encrypted = append(encrypted, char)
	}

	return strings.TrimSpace(string(encrypted))
}
