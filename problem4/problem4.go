// Find largest palindrome which is the product of two 3-digit integers.
package problem4

import (
	"strconv"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindrome(nombre int64) bool {
	chaine := strconv.FormatInt(nombre, 10)
	chaine_inverse := Reverse(chaine)
	return (chaine == chaine_inverse)
}

func FindLargestPalindrome() int64 {
	var max, mult, i, j int64 = 0, 0, 100, 100
	for i = 100; i < 1000; i = i + 1 {
		for j = 100; j < 1000; j = j + 1 {
			mult = i * j
			if IsPalindrome(mult) && mult > max {
				max = mult
			}
		}
	}
	return max
}
