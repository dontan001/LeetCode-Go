package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		strs := []string{"abc", "ab", "abcd"}
		prefix := LongestCommonPrefix(strs)
		t.Logf("the longest common prefix is: %s", prefix)
	})

	t.Run("case 2", func(t *testing.T) {
		strs := []string{"cir", "car"}
		prefix := LongestCommonPrefix(strs)
		t.Logf("the longest common prefix is: %s", prefix)
	})

	t.Run("case 3", func(t *testing.T) {
		strs := []string{"cir"}
		prefix := LongestCommonPrefix(strs)
		t.Logf("the longest common prefix is: %s", prefix)
	})
}
