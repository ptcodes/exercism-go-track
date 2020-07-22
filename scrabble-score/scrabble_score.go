package scrabble

import "strings"

func letterScore(letter rune) int {
	switch {
	case strings.ContainsRune("AEIOULNRST", letter):
		return 1
	case strings.ContainsRune("DG", letter):
		return 2
	case strings.ContainsRune("BCMP", letter):
		return 3
	case strings.ContainsRune("FHVWY", letter):
		return 4
	case strings.ContainsRune("K", letter):
		return 5
	case strings.ContainsRune("JX", letter):
		return 8
	case strings.ContainsRune("QZ", letter):
		return 10
	default:
		return 0
	}
}

func Score(input string) (score int) {
	input = strings.ToUpper(input)
	for _, letter := range input {
		score = score + letterScore(letter)
	}
	return
}
