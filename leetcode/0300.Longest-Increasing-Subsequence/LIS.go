package leetcode

func LongestIncreasingSubsequence(nums []int) int {

	var dp = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	var max int
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
