package letter

import "sync"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency returns the number of each unicode letter in a group of texts.
func ConcurrentFrequency(texts []string) FreqMap {
	frequencies := &concurrentFreqMap{freqMap: FreqMap{}}
	channel := make(chan *concurrentFreqMap)

	for _, text := range texts {
		go channelledFrequency(text, frequencies, channel)
	}

	for range texts {
		<-channel
	}

	return frequencies.freqMap
}

type concurrentFreqMap struct {
	sync.RWMutex
	freqMap FreqMap
}

func channelledFrequency(text string, frequencies *concurrentFreqMap, channel chan *concurrentFreqMap) {
	frequencies.Lock()
	defer frequencies.Unlock()

	for _, letter := range text {
		frequencies.freqMap[letter]++
	}

	channel <- frequencies
}
