package reverse

func Reverse(input string) (output string) {
	for _, runeValue := range input {
		output = string(runeValue) + output
	}
	return
}
