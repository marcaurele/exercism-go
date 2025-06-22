package wordcount

import (
	"regexp"
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	splits := regexp.MustCompile(`[\s,]+`).Split(strings.TrimSpace(phrase), -1)
	var frequencies Frequency = make(Frequency)
	for _, word := range splits {
		lower := ExtractWord(word)
		if lower == "" {
			continue
		}

		val, err := frequencies[lower]
		if !err {
			frequencies[lower] = 1
		} else {
			frequencies[lower] = val + 1
		}
	}
	return frequencies
}

func ExtractWord(input string) string {
	lower := strings.ToLower(input)
	var output []rune
	var prev rune
	for i, char := range lower {
		if prev == '\'' && (unicode.IsLetter(char) || unicode.IsDigit(char)) && i > 1 {
			output = append(output, '\'')
		}
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			output = append(output, char)
		}
		prev = char
	}
	return string(output)
}
