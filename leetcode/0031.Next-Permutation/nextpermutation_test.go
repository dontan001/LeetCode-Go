package leetcode

import "testing"

func TestNextPermutation(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		nums := []int{2, 3, 1}
		NextPermutation(nums)
		t.Logf("next permutation: %v", nums)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{2, 3, 1, 3, 3}
		NextPermutation(nums)
		t.Logf("next permutation: %v", nums)
	})
}
