package letter

import "sync"

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
	freqMap := FreqMap{}
	wg := sync.WaitGroup{}

	runeChannel := make(chan rune)

	for _, v := range l {
		wg.Add(1)

		go func(v string) {
			for _, r := range v {
				runeChannel <- r
			}

			wg.Done()
		}(v)
	}

	go func() {
		defer close(runeChannel)
		wg.Wait()
	}()

	for r := range runeChannel {
		freqMap[r]++
	}

	return freqMap
}
