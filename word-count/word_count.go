package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	frequencies := make(Frequency)
	for _, word := range regexp.MustCompile(`\w+('\w+)?`).FindAllString(strings.ToLower(phrase), -1) {
		frequencies[word]++
	}
	return frequencies
}
