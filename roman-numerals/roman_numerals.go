package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("input must be between 1 and 3999")
	}

	transform := func(in int, single string, middle string, max string) string {
		switch {
		case in <= 3:
			return strings.Repeat(single, in)
		case in == 4:
			return single + middle
		case in == 5:
			return middle
		case in <= 8:
			return middle + strings.Repeat(single, in-5)
		case in == 9:
			return single + max
		default:
			return ""
		}
	}

	output := transform(input/1000, "M", "", "")

	next := input % 1000
	output += transform(next/100, "C", "D", "M")

	next = next % 100
	output += transform(next/10, "X", "L", "C")

	u := next % 10
	output += transform(u, "I", "V", "X")

	return output, nil
}
