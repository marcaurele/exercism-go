package scrabble

import "strings"

func Score(word string) int {
	points := map[rune]int{
		'D': 2,
		'G': 2,
		'B': 3,
		'C': 3,
		'M': 3,
		'P': 3,
		'F': 4,
		'H': 4,
		'V': 4,
		'W': 4,
		'Y': 4,
		'K': 5,
		'J': 8,
		'X': 8,
		'Q': 10,
		'Z': 10,
	}

	score := 0
	for _, char := range strings.ToUpper(word) {
		if point, exists := points[char]; exists {
			score += point
		} else {
			score += 1
		}
	}
	return score
}
