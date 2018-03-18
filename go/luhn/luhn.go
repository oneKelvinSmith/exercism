package luhn

import (
	"errors"
	"unicode"
)

// Valid returns true if the given number is valid per the Luhn formula.
func Valid(input string) bool {
	number, err := filter(input)
	if len(number) <= 1 || err != nil {
		return false
	}

	return luhnSum(number)%10 == 0
}

func filter(input string) (number []int, err error) {
	var char rune
	for i := len(input) - 1; i >= 0; i-- {
		char = rune(input[i])

		if unicode.IsLetter(char) || unicode.IsPunct(char) || unicode.IsSymbol(char) {
			return nil, errors.New("invalid character present")
		}

		if unicode.IsDigit(char) {
			number = append(number, int(char-'0'))
		}
	}

	return
}

func luhnSum(number []int) (sum int) {
	var double int
	for index, digit := range number {
		if index%2 != 0 {
			double = digit * 2
			if double > 9 {
				double -= 9
			}
			digit = double
		}
		sum += digit
	}

	return
}
