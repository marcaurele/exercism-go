package atbash

import "strings"

func Atbash(s string) string {
	init := []rune(strings.ToLower(s))
	var result []rune
	letterSwitched := 0
	for i := 0; i < len(init); i++ {
		testLetter := init[i] >= 'a' && init[i] <= 'z'
		testNumber := init[i] >= '0' && init[i] <= '9'

		if testLetter || testNumber {
			if testLetter {
				result = append(result, ('z'-init[i])%26+'a')
			} else {
				result = append(result, init[i])
			}

			letterSwitched += 1
			if letterSwitched%5 == 0 {
				result = append(result, ' ')
			}
		}
	}
	return strings.TrimSpace(string(result))
}
