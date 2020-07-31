package romannumerals

import "errors"

var numbers = []struct {
	arabic int
	roman  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3000 {
		return "", errors.New("the number must be between 1 and 3000")
	}

	numeric := ""
	for _, v := range numbers {
		for n >= v.arabic {
			n -= v.arabic
			numeric += v.roman
		}
	}
	return numeric, nil
}
