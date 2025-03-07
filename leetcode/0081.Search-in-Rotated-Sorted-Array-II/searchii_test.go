package leetcode

import "testing"

func TestSearchII(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		nums := []int{2, 5, 6, 0, 0, 1, 2}
		target := 3

		exist := SearchII(nums, target)
		t.Log(exist)
	})
}
