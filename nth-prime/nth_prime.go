package prime

import "errors"

var primes = []int{}

func isPrime(n int) bool {
	for _, i := range primes {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("there is no zeroth prime")
	}

	for i := 2; ; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
		if len(primes) == n {
			break
		}
	}
	return primes[n-1], nil
}
