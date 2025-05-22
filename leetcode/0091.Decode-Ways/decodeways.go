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
