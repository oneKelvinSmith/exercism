package isogram

import (
	"sort"
	"unicode"
)

// IsIsogram determines if a given string is a word without repeating letters.
func IsIsogram(word string) bool {
	letters := filter(word)
	letters.reverseSort()

	for index := 0; index < len(letters)-1; index++ {
		if letters[index] == letters[index+1] {
			return false
		}
	}

	return true
}

func filter(word string) sortable {
	letters := []rune{}
	for _, letter := range word {
		if unicode.IsLetter(letter) {
			letters = append(letters, unicode.ToLower(letter))
		}
	}
	return sortable(letters)
}

func (letters sortable) reverseSort() {
	sort.Sort(letters)
}

type sortable []rune

func (letters sortable) Less(left, right int) bool {
	return letters[left] > letters[right]
}

func (letters sortable) Swap(left, right int) {
	letters[left], letters[right] = letters[right], letters[left]
}

func (letters sortable) Len() int {
	return len(letters)
}
