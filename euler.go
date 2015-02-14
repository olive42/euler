package main

import (
	"fmt"

	"github.com/olive42/euler/problem1"
	"github.com/olive42/euler/problem2"
	"github.com/olive42/euler/problem4"
	"github.com/olive42/euler/problem5"
)

func main() {
	fmt.Printf("Problem 1: %v\n", problem1.Problem1(1000))
	fmt.Printf("Problem 2: %v\n", problem2.FibSum(4000000))
	fmt.Printf("Problem 4: %v\n", problem4.FindLargestPalindrome())
	fmt.Printf("Problem 5: %v\n", problem5.FindSmallestMultiple(20))
}
