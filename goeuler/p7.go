package goeuler

import (
	util "euler/lib"
)

func nthPrime(n int) int {
	var result int = 2
	var primeNum int = 3
	for result < n {
		primeNum += 2
		if util.IsPrime(primeNum) {
			result++
		}
	}
	return primeNum
}

func P7() int {
	return nthPrime(10001)
}