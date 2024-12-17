package leetcode

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	combinations := [][]int{}

	var backtrack func([]int, int, int)
	backtrack = func(combination []int, remain int, start int) {
		if remain < 0 {
			return
		}

		if remain == 0 {
			found := []int{}
			found = append(found, combination...)
			combinations = append(combinations, found)
			return
		} else {
			var last int = -1
			for i := start; i < len(candidates); i++ {
				if last == candidates[i] {
					continue
				}

				if remain < candidates[i] {
					break
				}

				combination = append(combination, candidates[i])
				backtrack(combination, remain-candidates[i], i+1)

				last = combination[len(combination)-1]
				combination = combination[:len(combination)-1]
			}
		}
	}

	backtrack([]int{}, target, 0)
	return combinations
}
