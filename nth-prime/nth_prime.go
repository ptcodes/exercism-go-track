package prime

import "math"

func isPrime(n int) bool {
	for j := 3; j <= int(math.Sqrt(float64(n))); j += 2 {
		if n%j == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	if n > 2 && n%2 == 0 {
		return 0, false
	}
	counter := 0
	num := 2
	for {
		if isPrime(num) {
			counter++
			if counter == n {
				break
			}
		}
		num++
	}
	return num, true
}
