package problem2

import (
	"testing"
)

var fibtests = []struct {
	in  int
	out int
}{
	{0, 0},
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{5, 8},
	{6, 13},
	{7, 21},
	{8, 34},
}

var sumtests = []struct {
	in  int
	out int
}{
	{40, 44},
	{4000000, 4613732},
}

func testFibonacci(t *testing.T) {
	for _, tt := range fibtests {
		result := Fibonacci(tt.in)
		if tt.out != result {
			t.Error("Expected %v, got %v", tt.out, result)
		}
	}
}

func testFibsum(t *testing.T) {
	for _, tt := range sumtests {
		result := FibSum(tt.in)
		if tt.out != result {
			t.Error("Expected %v, got %v", tt.out, result)
		}
	}
}
