package leetcode

import "testing"

func TestThreeSumClosestBruteForce(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		nums := []int{-1, 2, 1, -4}
		target := 1
		sum := ThreeSumClosestBruteForce(nums, target)
		t.Log(sum)
	})
}
