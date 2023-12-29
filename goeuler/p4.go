package goeuler

import (
	"math"
	util "euler/lib"
)

// optimised solution , inspired from overview of P4
func largestPalindromeProduct(num int) int{
	largestPalindrome := 0
	a := int(math.Pow(10, float64(num))) - 1
	b, db := 1, 1
	c := int(math.Pow(10, float64(num) - 1.0))
	for a >= c {
		if a % 11 == 0 {
			b = int(math.Pow(10, float64(num))) - 1
			db = 1
		} else {
			b = int(math.Pow(10, float64(num))) - 10
			db = 11
		}

		for b >= a {
			if a * b <= largestPalindrome {
				break
			}
			if util.IsPalindrome(a * b) {
				largestPalindrome = a * b
			}
			b -= db
		}
		a--
	}
	return largestPalindrome
}

func P4() int{
	return largestPalindromeProduct(3)
}