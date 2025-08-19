package sieve

func Sieve(limit int) []int {
	allNumbers := make([]bool, limit+1)
	var primes []int
	for i := 2; i <= limit; i++ {
		if !allNumbers[i] {
			primes = append(primes, i)
			for j := 0; j*i <= limit; j++ {
				allNumbers[j*i] = true
			}
		}
	}
	return primes
}
