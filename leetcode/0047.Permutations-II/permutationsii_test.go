package leetcode

import "testing"

func TestPermuteUnique(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		input := []int{1, 1, 2}
		res := PermuteUnique(input)
		t.Logf("PermuteUnique %d : %v", len(res), res)
	})

	t.Run("case 2", func(t *testing.T) {
		input := []int{3, 3, 0, 3}
		res := PermuteUnique(input)
		t.Logf("PermuteUnique %d : %v", len(res), res)
	})
}
