package raindrops

import "strconv"

// Convert is fizzbuzz for the factors 3, 5 and 7.
func Convert(number int) string {
	raindrop := ""
	if number%3 == 0 {
		raindrop += "Pling"
	}

	if number%5 == 0 {
		raindrop += "Plang"
	}

	if number%7 == 0 {
		raindrop += "Plong"
	}

	if raindrop == "" {
		return strconv.Itoa(number)
	}

	return raindrop
}
