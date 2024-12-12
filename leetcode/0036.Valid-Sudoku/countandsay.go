package leetcode

import "strconv"

func CountAndSay(n int) string {

	var count func(string) string
	count = func(s string) string {
		if len(s) == 1 {
			return "11"
		}

		var result string = ""
		var total int = 1
		for i := 1; i <= len(s); i++ {
			if i == len(s) {
				result = result + strconv.Itoa(total) + string(s[i-1])
			} else {
				if s[i] == s[i-1] {
					total++
				} else {
					result = result + strconv.Itoa(total) + string(s[i-1])
					total = 1
				}
			}
		}
		return result
	}

	var say string
	for i := 1; i <= n; i++ {
		if i == 1 {
			say = "1"
		} else {
			say = count(say)
		}
	}
	return say
}
