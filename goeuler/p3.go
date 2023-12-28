package goeuler

func largePrimeFactor(num int) int {
	var last_fact int
	factor := 2
	if num % factor == 0 {
		num /= factor
		last_fact = factor
	}
	factor++
	for num > 1 {
		if num % factor == 0 {
			num /= factor
			last_fact = factor
			for num % factor == 0 {
				num /= factor
			}
		}
		factor = factor + 2
	}
	return last_fact
}

func P3() int {
	return largePrimeFactor(600851475143)
}