package problem1

// https://projecteuler.net/problem=1
// Sum all integers lower than limit that are multiple of 3 or 5.
func Problem1(limit int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
