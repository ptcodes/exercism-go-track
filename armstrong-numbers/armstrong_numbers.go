package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	number := strconv.Itoa(n)
	numberLength := float64(len(number))
	sum := 0.0
	for _, char := range number {
		digit := float64(char - '0')
		sum += math.Pow(digit, numberLength)
	}
	return n == int(sum)
}
