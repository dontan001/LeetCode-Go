package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		strs := []string{"abc", "ab", "abcd"}
		prefix := LongestCommonPrefix(strs)
		t.Logf("the longest common prefix is: %s", prefix)
	})
}
