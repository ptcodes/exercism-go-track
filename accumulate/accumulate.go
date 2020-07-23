package accumulate

type Converter func(string) string

func Accumulate(lines []string, fn Converter) (result []string) {
	for _, line := range lines {
		result = append(result, fn(line))
	}
	return
}
