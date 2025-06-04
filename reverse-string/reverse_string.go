package reverse

import (
	"unicode/utf8"
)

func Reverse(input string) string {
	result := []rune(input)
	for i, j := 0, utf8.RuneCountInString(input)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
