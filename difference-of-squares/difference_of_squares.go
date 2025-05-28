package diffsquares

func SquareOfSum(n int) int {
	total := (n * (n + 1) / 2)
	return total * total
}

func SumOfSquares(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += (i + 1) * (i + 1)
	}
	return total
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
