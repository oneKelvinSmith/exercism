package acronym

import (
	"unicode"
)

const (
	space  = rune(32)
	hyphen = rune(45)
)

// Abbreviate creates acronyms from a string.
func Abbreviate(s string) string {
	acronymRunes := []rune{}
	previousRune := space
	for _, r := range []rune(s) {
		if previousRune == space || previousRune == hyphen {
			acronymRunes = append(acronymRunes, unicode.ToUpper(r))
		}
		previousRune = r
	}

	return string(acronymRunes)
}
