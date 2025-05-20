package leetcode

import "sort"

func SubsetsWithDup(nums []int) [][]int {

	subsets := [][]int{}
	sort.Ints(nums)

	var sub func([]int, int)
	sub = func(prefix []int, start int) {
		subset := make([]int, 0, len(prefix))
		for i := 0; i < len(prefix); i++ {
			subset = append(subset, prefix[i])
		}
		subsets = append(subsets, subset)

		if start == len(nums) {
			return
		}

		var last int = -1000
		for i := start; i < len(nums); i++ {
			if nums[i] != last {
				prefix = append(prefix, nums[i])
				sub(prefix, i+1)
				prefix = prefix[:len(prefix)-1]
				last = nums[i]
			}
		}
	}

	sub([]int{}, 0)
	return subsets
}
