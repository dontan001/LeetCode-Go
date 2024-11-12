package leetcode

import "testing"

func TestThreeSum(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4}
		res := ThreeSum(nums)
		t.Logf("res:%v", res)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{0, 0, 0, 0}
		res := ThreeSum(nums)
		t.Logf("res:%v", res)
	})

	t.Run("case 3", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
		res := ThreeSum(nums)
		t.Logf("res:%v", res)
	})

	t.Run("case 4", func(t *testing.T) {
		nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
		res := ThreeSum(nums)
		t.Logf("res:%v", res)
	})
}
