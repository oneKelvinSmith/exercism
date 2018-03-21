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
	counting := make(chan FreqMap)

	for _, text := range texts {
		go func(text string) {
			counting <- Frequency(text)
		}(text)
	}

	frequencies = FreqMap{}
	finished := make(chan bool)
	go func() {
		for range texts {
			result := <-counting
			for letter, count := range result {
				frequencies[letter] += count
			}
		}
		finished <- true
	}()

	<-finished

	return
}
