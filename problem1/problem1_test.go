package problem1

import (
	"testing"
)

// https://projecteuler.net/problem=1 says for limit=10, result is 23
func testProblem1(t *testing.T) {
	expected := Problem1(10)
	if 23 != expected {
		t.Error("Expected 23, got ", expected)
	}
}
