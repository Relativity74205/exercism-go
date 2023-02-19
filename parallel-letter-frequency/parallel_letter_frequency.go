package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var wg sync.WaitGroup
	results := make([]FreqMap, len(l))
	wg.Add(len(l))
	for i, str := range l {
		str := str
		go func(i int) {
			results[i] = Frequency(str)
			wg.Done()
		}(i)
	}
	wg.Wait()

	finalResult := make(FreqMap)
	for _, freqMap := range results {
		for char, cnt := range freqMap {
			finalResult[char] += cnt
		}
	}
	return finalResult
}
