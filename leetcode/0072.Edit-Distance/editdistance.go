package leetcode

func minDistance(word1 string, word2 string) int {

	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	var m, n int = len(word1), len(word2)
	dp := make([][]int, m+1)
	for idx := 0; idx <= m; idx++ {
		dp[idx] = make([]int, n+1)
		dp[idx][0] = idx
	}

	for idx := 0; idx <= n; idx++ {
		dp[0][idx] = idx
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			c1, c2 := word1[i-1], word2[j-1]
			if c1 == c2 {
				dp[i][j] = dp[i-1][j-1]
			} else {
				replace := dp[i-1][j-1] + 1
				delete := dp[i-1][j] + 1
				insert := dp[i][j-1] + 1

				min := replace
				if delete < min {
					min = delete
				}
				if insert < min {
					min = insert
				}
				dp[i][j] = min
			}
		}
	}

	return dp[m][n]
}
