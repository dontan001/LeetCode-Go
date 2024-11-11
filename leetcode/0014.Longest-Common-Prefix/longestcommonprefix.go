package leetcode

func LongestCommonPrefix(strs []string) string {

	var prefix string = ""
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				break
			}

			if j == len(strs)-1 {
				prefix = prefix + string(strs[0][i])
			}
		}
	}

	return prefix
}
