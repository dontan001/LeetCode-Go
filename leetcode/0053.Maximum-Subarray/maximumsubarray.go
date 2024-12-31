package leetcode

func MaximumSubArray(nums []int) int {

	cache := make(map[int]int)
	var dp func(int) int
	dp = func(end int) int {
		if end == 0 {
			cache[end] = nums[0]
			return nums[0]
		}

		m, ok := cache[end-1]
		if !ok {
			m = dp(end - 1)
		}

		if m >= 0 {
			m = nums[end] + m
		} else {
			m = nums[end]
		}

		cache[end] = m
		return m
	}

	var max int = -2 << 31
	for i := 0; i < len(nums); i++ {
		t := dp(i)
		if t > max {
			max = t
		}
	}
	return max
}
