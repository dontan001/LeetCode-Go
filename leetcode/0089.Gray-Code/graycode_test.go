package leetcode

import "testing"

func TestGrayCode(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		t.Logf("2: %v", GrayCode(2))
		t.Logf("3: %v", GrayCode(3))
	})
}
