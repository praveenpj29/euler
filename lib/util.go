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

// Function to calculate the greatest common divisor (GCD)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the least common multiple (LCM)
func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return abs(a*b) / gcd(a, b)
}

// Function to find the absolute value
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Function to find the LCM of elements in an array
func LCMArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
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
