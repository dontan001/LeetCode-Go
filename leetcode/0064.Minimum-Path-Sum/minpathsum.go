package leetcode

func MinimumPathSum(grid [][]int) int {
	var m int = len(grid)
	var n int = len(grid[0])
	for i := 0; i < n; i++ {
		if i == 0 {
			continue
		}
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}

	for j := 0; j < m; j++ {
		if j == 0 {
			continue
		}
		grid[j][0] = grid[j-1][0] + grid[j][0]
	}

	var min func(int, int) int
	min = func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}

	return grid[m-1][n-1]
}
