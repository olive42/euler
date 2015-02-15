package problem12

import (
	"fmt"
)

const (
	LOG = true
)

func log(s string) {
	if LOG {
		fmt.Printf(s)
	}
}

// findNumDivisors returns the number of divisors for a given integer.
func findNumDivisors(op int) int {
	var divisors []int
	for i := 1; i <= op; i++ {
		if op%i == 0 {
			divisors = append(divisors, i)
		}
	}
	log(fmt.Sprintf("result is %v (%v).\n", len(divisors), divisors))
	return len(divisors)
}

func getTriangleNumber(i int) int {
	if i == 1 {
		return 1
	}
	return getTriangleNumber(i-1) + i
}

// Find500Divisors returns the first triangle number with 500 divisors
func Find500Divisors() int {
	i := 1
	for {
		t := getTriangleNumber(i)
		log(fmt.Sprintf("index: %v, triangle #: %v\n", i, t))
		if findNumDivisors(t) > 500 {
			return t
		}
		i++
	}
}
