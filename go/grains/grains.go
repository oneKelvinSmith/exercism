package grains

import "errors"

// Square returns the number of grains of wheat on a chess board for a given chess square.
func Square(position int) (number uint64, err error) {
	if position <= 0 || 64 < position {
		err = errors.New("invalid square")

		return
	}

	number = 1 << uint64(position-1)

	return
}

// Total returns the total amount of grain a chessboard represents.
func Total() uint64 {
	return 1<<64 - 1
}
