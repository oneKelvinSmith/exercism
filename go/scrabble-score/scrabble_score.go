package scrabble

import (
	"unicode"
)

var characterScores = map[rune]int{
	rune('a'): 1,
	rune('b'): 3,
	rune('c'): 3,
	rune('d'): 2,
	rune('e'): 1,
	rune('f'): 4,
	rune('g'): 2,
	rune('h'): 4,
	rune('i'): 1,
	rune('j'): 8,
	rune('k'): 5,
	rune('l'): 1,
	rune('m'): 3,
	rune('n'): 1,
	rune('o'): 1,
	rune('p'): 3,
	rune('q'): 10,
	rune('r'): 1,
	rune('s'): 1,
	rune('t'): 1,
	rune('u'): 1,
	rune('v'): 4,
	rune('w'): 4,
	rune('x'): 8,
	rune('y'): 4,
	rune('z'): 10,
}

// Score calculates an returns the scrabble score for a given word.
func Score(word string) int {
	score := 0
	for _, character := range word {
		score += scoreFor(character)
	}
	return score
}

func scoreFor(character rune) int {
	return characterScores[unicode.ToLower(character)]
}
