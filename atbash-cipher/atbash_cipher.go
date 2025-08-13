package atbash

import "strings"

func Atbash(s string) string {
	init := []rune(strings.ToLower(s))
	var result []rune
	letterSwitched := 0
	for i := 0; i < len(init); i++ {
		if init[i] >= 'a' && init[i] <= 'z' {
			result = append(result, ('z'-init[i])%26+'a')
			letterSwitched += 1
			if letterSwitched%5 == 0 {
				result = append(result, ' ')
			}
		} else if init[i] >= '0' && init[i] <= '9' {
			result = append(result, init[i])
			letterSwitched += 1
			if letterSwitched%5 == 0 {
				result = append(result, ' ')
			}
		}
	}
	return strings.TrimSpace(string(result))
}
