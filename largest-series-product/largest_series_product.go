package lsproduct

import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span == 0 {
		return 1, nil
	}
	if span < 0 {
		return -1, errors.New("span must not be negative")
	}
	if len(digits) < span {
		return -1, errors.New("span must be smaller than string length")
	}
	var max, product int64
	for i := 0; i <= len(digits)-span; i++ {
		product = 1
		for _, digit := range digits[i : i+span] {
			if !unicode.IsDigit(rune(digit)) {
				return -1, errors.New("digits input must only contain digits")
			}
			product *= int64(digit - '0')
		}
		if product > max {
			max = product
		}
	}
	return max, nil
}
