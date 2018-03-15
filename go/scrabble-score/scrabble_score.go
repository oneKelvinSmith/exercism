package scrabble

import (
	"unicode"
)

// Score calculates an returns the scrabble score for a given word.
func Score(word string) int {
	score := 0
	for _, letter := range word {
		score += valueFor(letter)
	}
	return score
}

func valueFor(letter rune) int {
	switch unicode.ToUpper(letter) {
	case 'Q', 'Z':
		return 10
	case 'J', 'X':
		return 8
	case 'K':
		return 5
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'B', 'C', 'M', 'P':
		return 3
	case 'D', 'G':
		return 2
	default:
		return 1
	}
}
