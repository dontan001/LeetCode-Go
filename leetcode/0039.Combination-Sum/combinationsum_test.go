package leetcode

import "testing"

func TestCombinationSum(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 10
		res := CombinationSum(candidates, target)
		t.Logf("res:%v", res)
	})

	t.Run("case 2 test slice", func(t *testing.T) {
		var doubleit func([]int)
		doubleit = func(nums []int) {
			for index, n := range nums {
				nums[index] = n * 2
			}
		}

		candidates := []int{1, 2, 3, 4}
		doubleit(candidates)
		t.Logf("res:%v", candidates) //[2 4 6 8]
	})

	t.Run("case 3 test array", func(t *testing.T) {
		var doubleit func([4]int)
		doubleit = func(nums [4]int) {
			for index, n := range nums {
				nums[index] = n * 2
			}
		}

		candidates := [4]int{}
		for i := 0; i < 4; i++ {
			candidates[i] = i + 1
		}
		doubleit(candidates)
		t.Logf("res:%v", candidates) //[1 2 3 4]
	})
}
