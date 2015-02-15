package problem10

import (
	"fmt"
)

const (
	// MAX = 10
	MAX = 2 * 1000 * 1000
	LOG = false
)

var isfilled = false

// Roughly 148k prime numbers to 2M
var primes = make([]int64, 150000)

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
	for i < int64(len(primes)) {
		p++
		if isPrime(p, i) {
			primes[i] = p
			log(fmt.Sprintf("%v is prime [index %v].\n", p, i))
			i++
		}
	}
	isfilled = true
}

// for unittests: for MAX = 10, sum = 17
func SumPrimes(max int64) int64 {
	if !isfilled {
		fillPrimes()
		log(fmt.Sprintf("largest prime: %v\n", primes[len(primes):]))
	}
	var sum int64 = 0
	for _, i := range primes {
		if i < MAX {
			sum = sum + i
		} else {
			break
		}
	}
	return sum
}
