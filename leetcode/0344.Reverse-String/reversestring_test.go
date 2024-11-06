package leetcode

import "testing"

func TestReverseString(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := []byte{'h', 'e', 'l', 'l', 'o'}
		ReverseString(input)
		t.Logf("ReverseString output: %v", string(input))
	})
}
