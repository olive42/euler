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
	// "github.com/olive42/euler/problem10" //slow
	"github.com/olive42/euler/problem11"
	// "github.com/olive42/euler/problem12" // very slow
	"github.com/olive42/euler/problem13"
	"github.com/olive42/euler/problem14"
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
	fmt.Printf("Problem 9: %v\n", a*b*c)

	// fmt.Printf("Problem 10: %v\n", problem10.SumPrimes(problem10.MAX))
	fmt.Printf("Problem 11: %v\n", problem11.FindMaxQuatruplet())
	// fmt.Printf("Problem 12: %v\n", problem12.Find500Divisors())
	fmt.Printf("Problem 13: %v\n", problem13.FindSum())
	fmt.Printf("Problem 14: %v\n", problem14.MaxChain())
}
