package leetcode

import "strconv"

func NumDecodings(s string) int {

	var num int
	var numdec func(string)
	numdec = func(sub string) {
		if len(sub) == 0 {
			num++
			return
		}

		number := sub[0] - '0'
		if number > 0 && number <= 9 {
			numdec(sub[1:])
		} else {
			return
		}

		if len(sub) > 1 {
			number, _ := strconv.Atoi(sub[:2])
			if number >= 10 && number <= 26 {
				numdec(sub[2:])
			} else {
				return
			}
		}
	}

	numdec(s)
	return num
}

func NumDecodings2(s string) int {

	var dp []int = make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		one := s[i] - '0'
		if one > 0 && one <= 9 {
			dp[i+1] += dp[i]
		}

		if i-1 >= 0 {
			two, _ := strconv.Atoi(s[i-1 : i+1])
			if two >= 10 && two <= 26 {
				dp[i+1] += dp[i-1]
			}
		}
	}

	return dp[len(s)]
}
