package leetcode

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n int = len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	var dp [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = 1
				if j > 0 && (dp[i][j-1] == 0 || obstacleGrid[i][j] == 1) {
					dp[i][j] = 0
				}
			} else if j == 0 {
				dp[i][j] = 1
				if dp[i-1][j] == 0 || obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				}
			} else {
				if obstacleGrid[i][j] == 0 {
					dp[i][j] = dp[i][j-1] + dp[i-1][j]
				} else {
					dp[i][j] = 0
				}
			}
		}
	}

	return dp[m-1][n-1]
}
