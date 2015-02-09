package problem5

import (
	"testing"
)

func testFactorial(t *testing.T) {
	if Factorial(10) != 3628800 {
		t.Error("Factorial doesn't work")
	}
}

func testIsMultipleOfEverything(t *testing.T) {
	if !IsMultipleOfEverything(24, 4) {
		t.Error("Multiple of everything doesn't work")
	}
}

func testFindSmallestMultiple(t *testing.T) {
	if FindSmallestMultiple(10) != 2520 {
		t.Error("Smallest multiple doesn't work")
	}
}
