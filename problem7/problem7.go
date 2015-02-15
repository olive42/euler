// Find 10001st prime number
// 2, 3, 5, 7, 11, and 13 -> 13 is 6th prime number
package problem7

import (
	"fmt"
)

const (
	TOP = 10001
	LOG = false
)

// Let's build an array of prime numbers
var primes = make([]int64, TOP)
var isfilled bool = false

func log(s string) {
	if LOG {
		fmt.Print(s)
	}
}

func isPrime(n int64, maxiter int64) bool {
	var i int64
	for i = 0; i < maxiter; i++ {
		if primes[i] == 0 {
			return false
		}
		if n%primes[i] == 0 {
			return false
		}
	}
	return true
}

func fillPrimes() {
	var p, i int64 = 2, 0
	primes[i] = p
	i++
	for i < TOP {
		p++
		if isPrime(p, i) {
			primes[i] = p
			log(fmt.Sprintf("%v is prime [index %v].\n", p, i))
			i++
		}
	}
	isfilled = true
}

func GetPrime(n int) int64 {
	if !isfilled {
		log(fmt.Sprintf("Computing primes...\n"))
		fillPrimes()
		log(fmt.Sprintf("Done.\n"))
	}
	return primes[n]
}
