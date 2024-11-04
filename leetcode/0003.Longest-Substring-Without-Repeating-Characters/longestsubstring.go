package leetcode

import (
	"fmt"
	"strings"
)

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

func LongestSubstringNoRepeatingV2(input string) int {
	var length, maxLength int = 0, 0
	var left, right int = 0, 0

	uniqueChars := make(map[string]bool, 0)
	for i := 0; i < len(input); i++ {
		currentChar := string(input[i])
		if _, ok := uniqueChars[currentChar]; !ok {
			uniqueChars[string(input[i])] = true
			length++
			if length > maxLength {
				maxLength = length
				right = i
			}
		} else {
			for j := left; j < i; j++ {
				if string(input[j]) != currentChar {
					delete(uniqueChars, string(input[j]))
					left++
					length--
				} else {
					left++
					break
				}
			}
		}
	}

	fmt.Printf("the substring of string %s is %s \n", input, input[right-maxLength+1:right+1])
	return maxLength
}
