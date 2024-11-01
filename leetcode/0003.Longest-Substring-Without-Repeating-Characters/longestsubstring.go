package leetcode

import "strings"

func LongestSubstringNoRepeating(input string) string {

	var longest string
	var stage string
	for i := 0; i < len(input); i++ {
		if !strings.Contains(stage, string(input[i])) {
			stage = stage + string(input[i])
		} else {
			if len(stage) > len(longest) {
				longest = stage
			}
			stage = string(input[i])
		}
	}

	if len(longest) == 0 || len(stage) > len(longest) {
		longest = stage
	}

	return longest
}
