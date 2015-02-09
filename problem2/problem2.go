package problem2

// This func is not used, it is just the standard recursive algorithm
// for Fibonacci; we want to reuse the results as we go in the
// second func and not compute the Fibonacci sequence every time.
func Fibonacci(start int) int {
	if start < 2 {
		return start
	} else {
		return Fibonacci(start-1) + Fibonacci(start-2)
	}
}

func FibSum(limit int) int {
	sum := 0
	for a, b := 0, 1; a < limit; a, b = b+a, a {
		if a%2 == 0 {
			sum += a
		}
	}
	return sum
}
