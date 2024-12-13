package leetcode

import "testing"

func TestCombinationSum(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 10
		res := CombinationSum(candidates, target)
		t.Logf("res:%v", res)
	})
}
