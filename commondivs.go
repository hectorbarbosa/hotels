package main

import (
	"fmt"
)

func CommonDivisors(s []int) ([]int, error) {
	if len(s) < 1 {
		err := fmt.Errorf("invalid input %v\n", s)
		return nil, err
	}

	if ContainsInt(s, 1) == true {
		err := fmt.Errorf("invalid input %v\n", s)
		return nil, err
	}

	divs, err := FindDivisors(s)
	if err != nil {
		return nil, err
	}
	divisors := divs[0]
	// fmt.Printf("%v\n", divisors)
	for i := 1; i < len(s); i++ {
		divisors = Intersection(divisors, divs[i])
		// fmt.Printf("%v\n", divisors)
	}

	return divisors, nil
}

func Intersection(s1 []int, s2 []int) []int {
	set := make([]int, 0)
	for _, num := range s1 {
		if ContainsInt(s2, num) == true {
			set = append(set, num)
		}
	}
	return set
}

func FindDivisors(s []int) ([][]int, error) {
	result := make([][]int, 0)
	for i, num := range s {
		if num < 1 {
			err := fmt.Errorf("invalid input %v\n", s)
			return nil, err
		}

		divs := make([]int, 0)
		for div := 2; div <= num; div++ {
			if remainder := s[i] % div; remainder == 0 {
				divs = append(divs, div)
			}
		}
		if len(divs) != 0 {
			result = append(result, divs)
		}
	}
	// fmt.Printf("%v\n", result)
	return result, nil
}

func ContainsInt(s []int, num int) bool {
	for _, n := range s {
		if n == num {
			return true
		}
	}
	return false
}
