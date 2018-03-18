package luhn

import (
	"strconv"
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
	for _, char := range input {
		if char == ' ' {
			continue
		}

		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}

		number = append(number, digit)
	}

	return
}

func luhnSum(number []int) (sum int) {
	for index, digit := range number {
		if (len(number)-index)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return
}
