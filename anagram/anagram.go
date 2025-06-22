package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subjectComp := SortString(subject)
	subjectLower := strings.ToLower(subject)
	var results []string

	for _, candidate := range candidates {
		if subjectComp == SortString(candidate) && subjectLower != strings.ToLower(candidate) {
			results = append(results, candidate)
		}
	}

	return results
}

func SortString(input string) string {
	small := strings.Split(strings.ToLower(input), "")
	sort.Strings(small)
	return strings.Join(small, "")
}
