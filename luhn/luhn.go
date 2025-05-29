package luhn

import (
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	lenId := len(id)
	if lenId <= 1 {
		return false
	}

	var total int
	moduloId := lenId % 2

	for pos, char := range id {
		val := int(char - '0')

		if val < 0 || val > 9 {
			return false
		}

		total += val
		if pos%2 == moduloId {
			total += val
			if val >= 5 {
				total -= 9
			}
		}

	}

	return total%10 == 0
}
