package leetcode

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var prefix string = ""
	var next bool = true
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				next = false
				break
			}

			if j == len(strs)-1 {
				prefix = prefix + string(strs[0][i])
			}
		}

		if !next {
			break
		}
	}

	return prefix
}
