package leetcode

import "testing"

func TestSearch(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 10

		idx := Search(nums, target)
		t.Logf("the index is %d", idx)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 1

		idx := Search(nums, target)
		t.Logf("the index is %d", idx)
	})

	t.Run("case 3", func(t *testing.T) {
		nums := []int{4, 5}
		target := 6

		idx := Search(nums, target)
		t.Logf("the index is %d", idx)
	})
}
