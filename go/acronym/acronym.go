package acronym

import (
	"unicode"
)

// Abbreviate creates acronyms from a string.
func Abbreviate(s string) string {
	acronymRunes := []rune{}
	previousRune := ' '
	for _, r := range s {
		if unicode.IsSpace(previousRune) || previousRune == '-' {
			acronymRunes = append(acronymRunes, unicode.ToUpper(r))
		}
		previousRune = r
	}

	return string(acronymRunes)
}
