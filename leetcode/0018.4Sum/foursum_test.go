package leetcode

import "testing"

func TestFourSum(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		nums := []int{2, 2, 2, 2, 2}
		res := FourSum(nums, 8)
		t.Log(res)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{-2, -1, -1, 1, 1, 2, 2}
		res := FourSum(nums, 0)
		t.Log(res)
	})
}
