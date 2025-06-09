package leetcode

import "testing"

func TestIsInterleave(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		s1 := "aabcc"
		s2 := "dbbca"
		s3 := "aadbbcbcac"
		expected := true
		actual := IsInterleave(s1, s2, s3)
		t.Logf("expected: %v, actual: %v", expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		s1 := "ab"
		s2 := "cd"
		s3 := "acbd"
		expected := true
		actual := IsInterleave(s1, s2, s3)
		t.Logf("expected: %v, actual: %v", expected, actual)
	})

	t.Run("case 3", func(t *testing.T) {
		s1 := "ab"
		s2 := "cd"
		s3 := "abdc"
		expected := false
		actual := IsInterleave(s1, s2, s3)
		t.Logf("expected: %v, actual: %v", expected, actual)
	})
}
