package leetcode

import "math"

func DivideTwoIntegers(dividend, divisor int) int {

	var positive bool
	if (dividend < 0 && divisor < 0) ||
		dividend >= 0 && divisor > 0 {
		positive = true
	}

	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	var quotient int
	dividend, divisor = abs(dividend), abs(divisor)
	for dividend >= divisor {
		var temp, power int = divisor, 0
		for temp <= dividend {
			power++
			temp = divisor << power
		}
		dividend = dividend - temp>>1
		quotient += 1 << (power - 1)
	}

	if quotient > math.MaxInt32 && positive {
		quotient = math.MaxInt32
	}

	if !positive {
		quotient = -quotient
	}
	return quotient
}
