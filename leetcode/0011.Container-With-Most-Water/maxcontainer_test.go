package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		max := MaxArea(height) // expected 49
		t.Logf("the max area is %d", max)
	})
}
