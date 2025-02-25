package leetcode

func Combine(n int, k int) [][]int {
	if k <= 0 {
		return [][]int{}
	}
	if k >= n {
		combination := []int{}
		for i := 1; i <= n; i++ {
			combination = append(combination, i)
		}
		return [][]int{combination}
	}

	combinations := [][]int{}
	var comb func([]int, int)
	comb = func(prefix []int, start int) {
		if len(prefix) == k {
			combination := []int{}
			combination = append(combination, prefix...)
			combinations = append(combinations, combination)
			return
		}

		if start > n {
			return
		}

		for i := start; i <= n; i++ {
			prefix = append(prefix, i)
			comb(prefix, i+1)
			prefix = prefix[:len(prefix)-1]
		}
	}

	comb([]int{}, 1)
	return combinations
}
