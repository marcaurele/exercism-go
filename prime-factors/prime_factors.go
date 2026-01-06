package prime

func Factors(n int64) []int64 {
	var primes = []int64{}

	if n < 1 {
		return []int64{}
	}

	var i int64

	for i = 2; i<= n; {
		if n%i == 0 {
			primes = append(primes, i)
			n /= i
		} else {
			i++
		}
		if n == 1 {
			break
		}
	}
	return primes
}
