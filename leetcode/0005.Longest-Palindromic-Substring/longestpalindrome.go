package leetcode

func LongestPalindrome(s string) string {
	expandAroundCenter := func(input string, left int, right int) string {
		for left >= 0 && right < len(input) && input[left] == input[right] {
			left--
			right++
		}

		return input[left+1 : right]
	}

	var longestPalindrome string
	for i := 0; i < len(s); i++ {
		odd := expandAroundCenter(s, i, i)
		even := expandAroundCenter(s, i, i+1)

		if len(odd) > len(longestPalindrome) {
			longestPalindrome = odd
		}

		if len(even) > len(longestPalindrome) {
			longestPalindrome = even
		}
	}

	return longestPalindrome
}
