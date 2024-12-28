package leetcode

func MyPow(x float64, n int) float64 {

	var result float64 = 1
	if n < 0 {
		n = -n
		x = 1 / x
	}

	for n > 0 {
		if n&1 == 1 {
			result *= x
		}

		x *= x
		n >>= 1
	}

	return result
}
