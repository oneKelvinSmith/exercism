package bob

import (
	"strings"
	"unicode"
)

// Remark is a convenience type for identifying what kind of remark it is.
type Remark struct {
	remark string
}

func (r Remark) isSilence() bool {
	return r.remark == ""
}

func (r Remark) isShouting() bool {
	hasLetters := strings.IndexFunc(r.remark, func(c rune) bool {
		return unicode.IsLetter(c)
	}) >= 0
	isUpcased := strings.ToUpper(r.remark) == r.remark

	return hasLetters && isUpcased
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

func (r Remark) isExasperated() bool {
	return r.isShouting() && r.isQuestion()
}

func newRemark(remark string) Remark {
	return Remark{strings.TrimSpace(remark)}
}

// Hey greets with an appropriate remark.
func Hey(remark string) string {
	r := newRemark(remark)
	switch {
	case r.isSilence():
		return "Fine. Be that way!"
	case r.isExasperated():
		return "Calm down, I know what I'm doing!"
	case r.isShouting():
		return "Whoa, chill out!"
	case r.isQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}
