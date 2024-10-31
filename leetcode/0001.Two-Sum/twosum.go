package leetcode

func TwoSumV1(integers []int, target int) []int {

	var i, j int
	for i = 0; i < len(integers); i++ {
		for j = i + 1; j < len(integers); j++ {
			if integers[i]+integers[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func TwoSumV2(integers []int, target int) []int {

	m := make(map[int]int, 0)
	for i := 0; i < len(integers); i++ {
		if _, ok := m[target-integers[i]]; ok {
			return []int{m[target-integers[i]], i}
		} else {
			m[integers[i]] = i
		}
	}

	return []int{}
}
