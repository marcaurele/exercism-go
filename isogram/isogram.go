package isogram

import "strings"

func IsIsogram(word string) bool {
	letters := map[rune]bool{}
	for _, letter := range strings.ToUpper(word) {
		if _, exists := letters[letter]; exists {
			return false
		} else if letter != '-' && letter != ' ' {
			letters[letter] = true
		}
	}
	return true
}
