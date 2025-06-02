package pangram

import "strings"

func IsPangram(input string) bool {
	letterSeen := map[rune]bool{}
	for _, char := range strings.ToLower(input) {
		if char >= 'a' && char <= 'z' {
			if _, exists := letterSeen[char]; exists {
				continue
			}
			letterSeen[char] = true
		}
	}
	return len(letterSeen) == 26
}
