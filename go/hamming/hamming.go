package hamming

import (
	"errors"
)

// Distance returns the number of differences between two homologous DNA strands,
// or an error if the strands are not of equal length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New(
			"please provide DNA strands of the same length",
		)
	}
	differences := 0
	for index := range a {
		if a[index] != b[index] {
			differences++
		}
	}
	return differences, nil
}
