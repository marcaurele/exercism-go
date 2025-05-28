package diffsquares

import "math"

func SquareOfSum(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i + 1
	}
	return int(math.Pow(float64(total), 2))
}

func SumOfSquares(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += int(math.Pow(float64(i+1), 2))
	}
	return total
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
