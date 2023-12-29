package goeuler

func sumSquareDiff(num int) int {
	n_sum := (num * (num + 1)) / 2
	n_square_sum := (n_sum * (2*num + 1)) / 3
	return (n_sum * n_sum) - n_square_sum
}

func P6() int {
	return sumSquareDiff(100)
}