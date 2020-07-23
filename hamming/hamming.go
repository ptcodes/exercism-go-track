package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands of unequal length")
	}

	differences := 0
	for i := range a {
		if a[i] != b[i] {
			differences += 1
		}
	}
	return differences, nil
}
