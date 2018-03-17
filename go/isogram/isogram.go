package isogram

import (
	"unicode"
)

// IsIsogram determines if a given string is a word without repeating letters.
func IsIsogram(word string) bool {
	letters := make(map[rune]int)

	for _, char := range word {
		if unicode.IsLetter(char) {
			letter := unicode.ToLower(char)

			if letters[letter] != 0 {
				return false
			}

			letters[letter]++
		}
	}

	return true
}
