package leetcode

import "sort"

func PermuteUnique(nums []int) [][]int {

	remain := func(used []int) []int {
		r := []int{}
		r = append(r, nums...)
		for len(used) > 0 {
			elem := used[0]
			var index int
			for i := 0; i < len(r); i++ {
				if r[i] == elem {
					index = i
					break
				}
			}

			if index == 0 {
				r = r[1:]
			} else if index == len(r)-1 {
				r = r[:len(r)-1]
			} else {
				r = append(r[:index], r[index+1:]...)
			}
			used = used[1:]
		}
		return r
	}

	sort.Ints(nums)
	var permutations [][]int
	var dopermute func([]int)
	dopermute = func(permutation []int) {
		if len(permutation) == len(nums) {
			found := []int{}
			found = append(found, permutation...)
			permutations = append(permutations, found)
			return
		}

		remaining := remain(permutation)
		var last int = 100
		for i := 0; i < len(remaining); i++ {
			if remaining[i] == last {
				continue
			}
			permutation = append(permutation, remaining[i])
			dopermute(permutation)

			last = permutation[len(permutation)-1]
			permutation = permutation[:len(permutation)-1]
		}
	}

	dopermute([]int{})
	return permutations
}
