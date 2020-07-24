package pangram

import (
	"strings"
)

func IsPangram(phrase string) bool {
	if phrase == "" {
		return false
	}
	phrase = strings.ToLower(phrase)

	letters := make(map[rune]bool)
	for _, char := range phrase {
		if char >= 'a' && char <= 'z' {
			letters[char] = true
		}
	}
	return len(letters) == 26
}
