package isbn

import "regexp"

func IsValidISBN(isbn string) bool {
	isbn = regexp.MustCompile(`[^\dX]`).ReplaceAllString(isbn, "")
	if len(isbn) != 10 {
		return false
	}
	if !regexp.MustCompile(`[\d]{9,10}X?`).MatchString(isbn) {
		return false
	}
	sum := 0
	for i, digit := range isbn {
		if digit == 'X' {
			sum += 10
		} else {
			sum += (int(digit) - '0') * (10 - i)
		}
	}
	return sum%11 == 0
}
