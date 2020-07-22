package pascal

func factorial(n int) int {
	if n > 0 {
		result := n * factorial(n-1)
		return result
	}
	return 1
}

func formula(row, term int) int {
	return factorial(row) / (factorial(term) * factorial(row-term))
}

func Triangle(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		row := make([]int, 0)
		for j := 0; j <= i; j++ {
			value := formula(i, j)
			row = append(row, value)
		}
		result = append(result, [][]int{row}...)
	}
	return result
}
