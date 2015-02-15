package problem9

// find triplets a^2 + b^2 = c^2; a < b < c
// find the one where a+b+c == 1000
// return a*b*c
// Sqrt(1000) = 32, so let's not get too far

import (
	"fmt"
)

const (
	LOG = true
)

func log(s string) {
	if LOG {
		fmt.Print(s)
	}
}

func FindGoldenTriplet() (int, int, int) {
	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			for c := b + 1; c < 1000; c++ {
				//log(fmt.Sprintf("Operating on %v, %v, %v\n", a, b, c))
				if a*a+b*b == c*c {
					log(fmt.Sprintf("Found triplet: %v, %v, %v\n", a, b, c))
					if a+b+c == 1000 {
						log(fmt.Sprintf("And this one was golden!\n"))
						return a, b, c
					}
				}
			}
		}
	}
	return 0, 0, 0
}
