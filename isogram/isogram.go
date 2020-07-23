package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, "-", "")

	chars := make(map[rune]bool)
	for _, char := range strings.ToLower(word) {
		chars[char] = true
	}
	return len(word) == len(chars)
}
