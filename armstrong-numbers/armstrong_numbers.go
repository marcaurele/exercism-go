package armstrong

import (
	"math"
)

func IsNumber(n int) bool {
	var parts []int
	number := n
	for {
		if number < 10 {
			parts = append(parts, number%10)
			break
		}
		parts = append(parts, number%10)
		number = number / 10
	}
	l := len(parts)
	total := 0.0
	for i := 0; i < l; i++ {
		total += math.Pow(float64(parts[i]), float64(l))
	}
	return int(total) == n
}
