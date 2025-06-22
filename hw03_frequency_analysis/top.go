package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strings"
)

type frequency struct {
	word  string
	count int
}

func (f frequency) String() string {
	return fmt.Sprintf("%s - %d", f.word, f.count)
}

var re = regexp.MustCompile(`[[:punct:]]`)

func Top10(s string) []string {
	freqWords := countWords(clearText(s))

	top := 10
	resp := make([]string, 0, top)
	for i := 0; i < top && i < len(freqWords); i++ {
		resp = append(resp, freqWords[i].word)
	}

	return resp
}

func clearText(s string) []string {
	return strings.Fields(re.ReplaceAllString(s, ""))
}

func countWords(words []string) []frequency {
	freqWords := []frequency{}

	for _, word := range words {
		word = strings.ToLower(word)

		if idx := slices.IndexFunc(freqWords, func(f frequency) bool {
			return f.word == word
		}); idx != -1 {
			// word exists
			freqWords[idx].count++
			continue
		}
		// new word
		freqWords = append(freqWords, frequency{word: word, count: 1})
	}

	sortFreqWords(freqWords)

	return freqWords
}

func sortFreqWords(freqWords []frequency) {
	sort.Slice(freqWords, func(i, j int) bool {
		if freqWords[i].count == freqWords[j].count {
			return freqWords[i].word < freqWords[j].word
		}
		return freqWords[i].count > freqWords[j].count
	})
}
