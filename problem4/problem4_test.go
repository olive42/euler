package problem4

import (
	"testing"
)

var reversetests = []struct {
	in  string
	out string
}{
	{"test", "tset"},
	{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
}

var palindrometests = []struct {
	in  int64
	out bool
}{
	{9009, true},
	{12345678987654321, true},
	{123456789, false},
}

func testReverse(t *testing.T) {
	for _, tt := range reversetests {
		result := Reverse(tt.in)
		if tt.out != result {
			t.Error("Expected %v, got %v", tt.out, result)
		}
	}
}

func testIsPalindrome(t *testing.T) {
	for _, tt := range palindrometests {
		ispal := IsPalindrome(tt.in)
		if tt.out != ispal {
			t.Error("Expected %v to be %v, turned out %v instead", tt.in, tt.out, ispal)
		}
	}
}
