package goeuler

import (
	util "euler/lib"
)

func primeSummation(num int) int {
	sum := 10 // starting after 3 primes , 2 + 3 + 5 = 10
	counter := 7
	for counter < num {
		if util.IsPrime(counter) {
			sum += counter
		}
		counter += 2
	}
	return sum
}

func (e *EulerFuncs) P10() int {
	return primeSummation(2000000)
}