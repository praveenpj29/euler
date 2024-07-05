package goeuler

import (
	util "euler/lib"
)

func summation(num, lim int) int {
	lim = lim - 1
	mod := lim / num
	result := ((num * mod * mod) + (num * mod)) / 2
	return result
}

func cal(nums []int, limit int) int {
	// greatest common factor of numbers present in nums array
	lcm := util.LCMArray(nums)
	var result int = 0
	for _, v := range nums {
		result += summation(v, limit)
	}
	return result - summation(lcm, limit)
}

func (e *EulerFuncs) P1() int {
	return cal([]int{3, 5}, 1000)
}