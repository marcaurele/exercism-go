package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("there is no zeroth prime")
	}

	primes := []int{}
	for i := 2; ; i++ {
		prime := true

		for _, val := range primes {
			if i%val == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes = append(primes, i)
		}

		if len(primes) == n {
			break
		}
	}
	return primes[n-1], nil
}
