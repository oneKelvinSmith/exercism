package raindrops

import "strconv"

// Convert is fizzbuzz for the factors 3, 5 and 7.
func Convert(number int) string {
	raindrop := ""
	if factorOf(number, 3) {
		raindrop += "Pling"
	}

	if factorOf(number, 5) {
		raindrop += "Plang"
	}

	if factorOf(number, 7) {
		raindrop += "Plong"
	}

	if raindrop == "" {
		return strconv.Itoa(number)
	}

	return raindrop
}

func factorOf(number, factor int) bool {
	return number%factor == 0
}
