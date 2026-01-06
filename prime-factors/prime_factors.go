package prime

func Factors(n int64) []int64 {
	var primes = []int64{}

	if n < 1 {
		return []int64{}
	}

	var i int64

	for i = 2; ; i++ {
		if n%i == 0 {
			primes = append(primes, i)
			n = n / i
			i = 1
		}
		if n == 1 {
			break
		}
	}
	return primes
}
