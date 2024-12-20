package leetcode

// JumpGameII Time Limit Exceeded
func JumpGameII(nums []int) int {

	var minSteps int = 10000
	cache := make(map[int]int)

	var backtrack func(int, int)
	backtrack = func(steps int, from int) {
		if from >= len(nums)-1 {
			if steps < minSteps {
				minSteps = steps
			}
			return
		}

		max := nums[from]
		if max == 0 {
			return
		}

		for i := 1; i <= max; i++ {
			if v, ok := cache[from+i]; ok && v < steps+1 {
				continue
			}

			cache[from+i] = steps + 1
			backtrack(steps+1, from+i)
		}
	}

	backtrack(0, 0)
	return minSteps
}

// JumpGameIIGreedy Accepted
func JumpGameIIGreedy(nums []int) int {

	var next, jumps int = 0, 0
	for next < len(nums)-1 {
		max := nums[next]
		current, farthest := next, 0
		for i := 1; i <= max; i++ {
			if current+i < len(nums)-1 {
				if current+i+nums[current+i] >= farthest {
					farthest = current + i + nums[current+i]
					next = current + i
				}
			} else {
				next = len(nums) - 1
				break
			}
		}
		jumps++
	}

	return jumps
}
