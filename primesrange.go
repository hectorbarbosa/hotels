package main

import (
	"fmt"
)

func GetPrimesRange(s []int) ([]int, error) {
	var primes = make([]int, 0)

	// invalid slice length, returning error
	if len(s) > 2 || len(s) == 0 {
		err := fmt.Errorf("invalid input %v\n", s)
		return nil, err
	}
	// second number is less then first, returning error
	if len(s) == 2 {
		if s[1] < s[0] {
			err := fmt.Errorf("invalid input %v\n", s)
			return nil, err
		}
		if s[0] < 2 {
			err := fmt.Errorf("invalid input %v\n", s)
			return nil, err
		}
	}

	for v := s[0]; v <= s[len(s)-1]; v++ {
		if isPrime(v) == true {
			primes = append(primes, v)
		}
	}
	return primes, nil
}

func isPrime(n int) bool {
	// 1 is not a prime number
	if n == 1 {
		return false
	}

	var isprime = true
	for i := 2; i < n; i++ {
		rem := n % i
		if rem == 0 {
			isprime = false
		}
	}
	return isprime
}
