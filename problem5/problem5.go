package problem5

func IsMultipleOfEverything(target int, upto int) bool {
	for i := 1; i <= upto; i++ {
		if target%i != 0 {
			return false
		}
	}
	return true
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func FindSmallestMultiple(upto int) int {
	for i := 1; i < Factorial(upto); i++ {
		if IsMultipleOfEverything(i, upto) {
			return i
		}
	}
	return 0
}
