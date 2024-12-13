package leetcode

import "sort"

func CombinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)

	combinations := [][]int{}
	var backtrack func([]int, int, int)
	backtrack = func(combination []int, remain int, start int) {
		if remain < 0 {
			return
		} else if remain == 0 {
			found := []int{}
			found = append(found, combination...)
			combinations = append(combinations, found)
		} else {
			for i := start; i < len(candidates); i++ {
				if candidates[i] > remain {
					break
				}

				combination = append(combination, candidates[i])
				backtrack(combination, remain-candidates[i], i)
				combination = combination[:len(combination)-1]
			}
		}
	}

	backtrack([]int{}, target, 0)
	return combinations
}
