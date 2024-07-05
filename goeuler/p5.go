package goeuler

import (
	util "euler/lib"
)

func smallestMultiple(len int) int {
	result := 0
	num_arr := []int{}
	for i:= 1 ; i <= len ; i++ {
		num_arr = append(num_arr, i)
	}
	result = util.LCMArray(num_arr)
	return result
}

func (e *EulerFuncs) P5() int {
	return smallestMultiple(20)
}