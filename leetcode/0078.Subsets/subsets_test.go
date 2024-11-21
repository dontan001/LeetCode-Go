package leetcode

import "testing"

func TestSubsetMask(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		nums := []int{}
		subsets := SubsetMask(nums)
		t.Logf("subsets: %v", subsets)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{0}
		subsets := SubsetMask(nums)
		t.Logf("subsets: %v", subsets)
	})

	t.Run("case 3", func(t *testing.T) {
		nums := []int{0, 1}
		subsets := SubsetMask(nums)
		t.Logf("subsets: %v", subsets)
	})
}
