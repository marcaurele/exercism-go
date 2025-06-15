package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	check := 0
	compactIsbn := strings.TrimSpace(strings.ReplaceAll(isbn, "-", ""))

	if len(compactIsbn) != 10 {
		return false
	}

	for i, char := range compactIsbn {
		if !(char >= '0' && char <= '9' || char == 'X' && i == 9) {
			return false
		}

		if char == 'X' {
			check += 10
		} else {
			check += (int(char) - '0') * (10 - i)
		}
	}

	return check%11 == 0
}
