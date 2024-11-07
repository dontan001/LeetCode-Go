package leetcode

import "strings"

func ZigZagConvert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	heads := make([][]string, numRows)

	var index, direction int = 0, 1
	for i := 0; i < len(s); i++ {
		heads[index] = append(heads[index], string(s[i]))

		if index == numRows-1 {
			direction = -1
		} else if index == 0 {
			direction = 1
		}
		index = index + direction
	}

	var result string
	for i := 0; i < numRows; i++ {
		result = result + strings.Join(heads[i], "")
	}
	return result
}
