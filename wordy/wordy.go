package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	if regexp.MustCompile(`\d \d`).MatchString(question) {
		return 0, false
	}
	operations := regexp.MustCompile(`(plus|minus|multiplied|divided)`).FindAllString(question, -1)
	numbers := regexp.MustCompile(`\-?\d+`).FindAllString(question, -1)
	if len(operations) == 0 && !regexp.MustCompile(`What is \d+\?`).MatchString(question) {
		return 0, false
	}
	if len(operations) != len(numbers)-1 {
		return 0, false
	}
	sum, _ := strconv.Atoi(numbers[0])
	for i, operation := range operations {
		num, _ := strconv.Atoi(numbers[i+1])
		switch operation {
		case "plus":
			sum += num
		case "minus":
			sum -= num
		case "multiplied":
			sum *= num
		case "divided":
			sum /= num
		}
	}
	return sum, true
}
