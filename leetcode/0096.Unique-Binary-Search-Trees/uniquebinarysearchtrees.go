package leetcode

func NumTrees(n int) int {
	if n <= 2 {
		return n
	}

	var countBSTrees func(int, int) int
	countBSTrees = func(start, end int) int {
		var temp int
		if start > end {
			return 1
		}

		for i := start; i <= end; i++ {
			left := countBSTrees(start, i-1)
			right := countBSTrees(i+1, end)
			temp += left * right
		}

		return temp
	}

	return countBSTrees(1, n)
}

func NumTreesDP(n int) int {
	var dp []int = make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
