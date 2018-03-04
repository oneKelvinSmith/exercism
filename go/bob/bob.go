package bob

import (
	"strings"
	"unicode"
)

const (
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Remark is a convenience type for identifying what kind of remark it is.
type Remark struct {
	remark string
}

func (r Remark) isSilence() bool {
	return strings.TrimSpace(r.remark) == ""
}

func (r Remark) isShouting() bool {
	trimmed := strings.TrimFunc(r.remark, func(char rune) bool {
		return unicode.IsDigit(char) ||
			unicode.IsPunct(char) ||
			unicode.IsSpace(char)
	})

	return strings.ToUpper(trimmed) == trimmed &&
		strings.ContainsAny(r.remark, upper)
}

func (r Remark) isQuestion() bool {
	trimmed := strings.TrimSpace(r.remark)
	return strings.HasSuffix(trimmed, "?")
}

// Hey greets with an appropriate remark.
func Hey(remark string) string {
	r := Remark{remark}
	switch {
	case r.isSilence():
		return "Fine. Be that way!"
	case r.isShouting():
		if r.isQuestion() {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case r.isQuestion():
		return "Sure."

	default:
		return "Whatever."
	}
}
