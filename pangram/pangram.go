package pangram

import (
	"unicode"
)

func IsPangram(input string) bool {
	letterSeen := map[rune]bool{}
	for _, char := range input {
		lowerChar := unicode.ToLower(char)
		if _, exists := letterSeen[lowerChar]; lowerChar >= 'a' && lowerChar <= 'z' && !exists {
			letterSeen[char] = true
		}
	}
	return len(letterSeen) == 26
}
