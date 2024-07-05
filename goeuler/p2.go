package goeuler

func evenFibSum(num int) int{
	x , y := 1, 1
	sum := 0
	for x < num && y < num {
		sum += x + y
		x, y = (x + 2 * y), (2 * x + 3 * y)
	}
	return sum
}

func (e *EulerFuncs) P2() int {
	return evenFibSum(4000000)
}

