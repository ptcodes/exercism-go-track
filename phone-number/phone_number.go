package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func Number(number string) (string, error) {
	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")

	re := regexp.MustCompile(`^1?[2-9]{1}[0-9]{2}[2-9]{1}[0-9]{6}$`)
	if !re.MatchString(number) {
		return "", errors.New("incorrect format")
	}

	if len(number) == 11 {
		number = number[1:]
	}
	return number, nil
}

func AreaCode(number string) (string, error) {
	number, err := Number(number)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(number string) (string, error) {
	number, err := Number(number)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:]), nil
}
