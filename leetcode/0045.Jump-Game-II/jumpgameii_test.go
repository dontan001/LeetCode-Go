package leetcode

import "testing"

func TestJumpGameII(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		min := JumpGameIIGreedy(nums)
		t.Logf("min steps: %d", min)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{3, 2, 1}
		min := JumpGameIIGreedy(nums)
		t.Logf("min steps: %d", min)
	})

	t.Run("case 3", func(t *testing.T) {
		nums := []int{2, 3, 1}
		min := JumpGameIIGreedy(nums)
		t.Logf("min steps: %d", min)
	})
}
