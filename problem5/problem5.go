// Find the smallest integer evenly divisible by all numbers from 1 to 20.
package problem5

// IsMultipleOfEverything will return true if it is evenly divisible
// by all integers from 1 to upto.
func IsMultipleOfEverything(target int, upto int) bool {
	for i := 1; i <= upto; i++ {
		if target%i != 0 {
			return false
		}
	}
	return true
}

// Factorial computes factorial(n).
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// FindSmallestMultiple finds the smallest integer evenly divisible by
// all numbers from 1 to upto.
func FindSmallestMultiple(upto int) int {
	for i := 1; i < Factorial(upto); i++ {
		if IsMultipleOfEverything(i, upto) {
			return i
		}
	}
	return 0
}
