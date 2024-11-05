package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := "abc"
		output := LongestPalindrome(input)
		t.Logf("output of input %s : %v", input, output)
	})

	t.Run("case 2", func(t *testing.T) {
		input := "abbabb"
		output := LongestPalindrome(input)
		t.Logf("output of input %s : %v", input, output)
	})
}
