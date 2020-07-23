package luhn

import (
	"strings"
	"unicode"
)

func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}

	if len(input)%2 == 1 {
		input = "0" + input
	}

	sum := 0
	for i, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
		number := int(char - '0')
		if i%2 == 0 {
			number = number * 2
			if number > 9 {
				number = number - 9
			}
		}
		sum = sum + number
	}

	return sum%10 == 0
}
