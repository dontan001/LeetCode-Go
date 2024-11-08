package leetcode

func ReverseInteger(x int) int {
	t := 0
	for x != 0 {
		t = t*10 + x%10
		x = x / 10
	}

	if t > 1<<31-1 || t < -1<<31 {
		t = 0
	}

	return t
}
