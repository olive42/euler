package main

import (
	"fmt"

	// "github.com/olive42/euler/problem1"
	// "github.com/olive42/euler/problem2"
	// "github.com/olive42/euler/problem4"
	// "github.com/olive42/euler/problem5"
	// "github.com/olive42/euler/problem6"
	// "github.com/olive42/euler/problem7"
	"github.com/olive42/euler/problem8"
	"github.com/olive42/euler/problem9"
)

func main() {
	// fmt.Printf("Problem 1: %v\n", problem1.Problem1(1000))
	// fmt.Printf("Problem 2: %v\n", problem2.FibSum(4000000))
	// fmt.Printf("Problem 4: %v\n", problem4.FindLargestPalindrome())
	// fmt.Printf("Problem 5: %v\n", problem5.FindSmallestMultiple(20))
	// fmt.Printf("Problem 6: %v\n", problem6.DiffSquaresSums(100))
	// fmt.Printf("Problem 7: %v\n", problem7.GetPrime(10000))

	max, err := problem8.FindLargestMultInString(problem8.NUMBER)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Problem 8: %v\n", max)

	a, b, c := problem9.FindGoldenTriplet()
	fmt.Printf("Problem9: %v\n", a*b*c)
}
