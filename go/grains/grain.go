package grains

import "errors"

// Square returns the number of grains of wheat on a chess board for a given chess square.
func Square(number int) (count uint64, err error) {
	if number <= 0 || 64 < number {
		err = errors.New("invalid square")

		return
	}

	count = square(number)

	return
}

func square(number int) uint64 {
	if number == 1 {
		return 1
	}

	return 2 * square(number-1)

}

// Total returns the total amount of grain a chessboard represents.
func Total() uint64 {
	return 18446744073709551615
}
