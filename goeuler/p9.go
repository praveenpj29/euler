package goeuler

// Not an optimised solution.
func (e *EulerFuncs) P9() int {
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			for k := 1; k < 1000; k++ {
				if i + j + k == 1000 {
					if (i * i) + (j * j) == (k * k) {
						return i*j*k
					}
				}
			}
		}
	}
	return 1
}