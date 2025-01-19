package leetcode

func MyAtoi(s string) int {

	var digitStart bool
	var positive bool = true
	var num int
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !digitStart {
			if c == ' ' {
				continue
			} else {
				digitStart = true
				if c == '-' {
					positive = false
					continue
				} else if c == '+' {
					continue
				}
			}
		}

		if c < '0' || c > '9' {
			break
		} else {
			num = num*10 + int(c-'0')
			if positive {
				if num >= 1<<31-1 {
					num = 1<<31 - 1
					break
				}
			} else {
				if num >= 1<<31 {
					num = 1 << 31
					break
				}
			}
		}
	}

	if !positive {
		num *= -1
	}

	return num
}
