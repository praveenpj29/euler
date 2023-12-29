package util

import (
	"math"
)

func IsPrime(num int) bool {
	if num <= 3 {
		return true
	}
	for i := 3; i < int(math.Sqrt(float64(num))) ; i = i + 2 {
		if num % i == 0 {
			return false
		}
	}
	return true
}

// Function to find the LCM of elements in an array
// Euclidean algorithm
func LCMArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	result := arr[0] 
	for i:= 1 ; i < len(arr); i++ {
		num1 := result
		num2 := arr[i]
		gcd := 1
		for num2 > 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		result = (result * arr[i]) / gcd
	}
	return result
}

func Reverse(num int) int{
	reversed := 0
	for num > 0 {
		reversed = 10 * reversed + num % 10
		num /= 10
	}
	return reversed
} 

func IsPalindrome(num int) bool {
	return num == Reverse(num)
}
