package leetcode

import "testing"

func TestLongestSubstring(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := "abcabcbb"
		sub := LongestSubstringNoRepeating(input)
		t.Logf("The substring of %s is %s", input, sub)
	})

	t.Run("case 2", func(t *testing.T) {
		input := "bbbbb"
		sub := LongestSubstringNoRepeating(input)
		t.Logf("The substring of %s is %s", input, sub)
	})

	t.Run("case 3", func(t *testing.T) {
		input := "pwwkew"
		sub := LongestSubstringNoRepeating(input)
		t.Logf("The substring of %s is %s", input, sub)
	})

	t.Run("case 4", func(t *testing.T) {
		input := " "
		sub := LongestSubstringNoRepeating(input)
		t.Logf("The substring of%sis%s...", input, sub)
	})

	t.Run("case 5", func(t *testing.T) {
		input := "abcdef"
		sub := LongestSubstringNoRepeating(input)
		t.Logf("The substring of %s is %s", input, sub)
	})

	t.Run("case 6", func(t *testing.T) {
		input := "aab"
		sub := LongestSubstringNoRepeating(input)
		t.Logf("The substring of %s is %s", input, sub)
	})
}

func TestLongestSubstringNoRepeatingV2(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := "dvdf"
		max := LongestSubstringNoRepeatingV2(input)
		t.Logf("The max length of substring is %d", max)
	})

	t.Run("case 2", func(t *testing.T) {
		input := "abcdefcghifmj"
		max := LongestSubstringNoRepeatingV2(input)
		t.Logf("The max length of substring is %d", max)
	})
}
