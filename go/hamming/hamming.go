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
	for index := 0; index < len(a); index++ {
		if []rune(a)[index] != []rune(b)[index] {
			differences++
		}
	}
	return differences, nil
}
