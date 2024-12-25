package leetcode

func Permute(nums []int) [][]int {

	remain := func(used []int) []int {
		res := []int{}
		usedMap := make(map[int]bool)
		if len(used) > 0 {
			for _, u := range used {
				usedMap[u] = true
			}
		}

		for _, n := range nums {
			if _, ok := usedMap[n]; ok {
				continue
			} else {
				res = append(res, n)
			}
		}

		return res
	}

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
		for i := 0; i < len(remaining); i++ {
			permutation = append(permutation, remaining[i])
			dopermute(permutation)
			permutation = permutation[:len(permutation)-1]
		}
	}

	dopermute([]int{})
	return permutations
}
