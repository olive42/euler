package problem6

func sumSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i*i
	}
	return sum
}

func squareSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum * sum
}

func DiffSquaresSums(n int) int {
	return squareSum(n) - sumSquares(n)
}
