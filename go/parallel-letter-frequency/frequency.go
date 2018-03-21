package letter

// FreqMap represents the frequencies of letters.
type FreqMap map[rune]int

// Frequency returns the number of occurances of each unicode letter in a string.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency returns the number of each unicode letter in a group of texts.
func ConcurrentFrequency(texts []string) (frequencies FreqMap) {
	channel := make(chan FreqMap)
	defer close(channel)

	for _, text := range texts {
		go func(text string) {
			channel <- Frequency(text)
		}(text)
	}

	frequencies = FreqMap{}
	for range texts {
		result := <-channel
		for letter := range result {
			frequencies[letter] += result[letter]
		}
	}

	return
}
