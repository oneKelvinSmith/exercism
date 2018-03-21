package letter

// ConcurrentFrequency returns the number of each unicode letter in a group of texts.
func ConcurrentFrequency(texts []string) (frequencies FreqMap) {
	frequencies = make(FreqMap)
	results := make(chan FreqMap, len(texts))

	for _, text := range texts {
		go channelledFrequency(text, results)
	}

	var channelFrequencies FreqMap
	for range texts {
		channelFrequencies = <-results

		for letter := range channelFrequencies {
			frequencies[letter] += channelFrequencies[letter]
		}
	}

	return
}

func channelledFrequency(text string, results chan FreqMap) {
	results <- Frequency(text)
}
