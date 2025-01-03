package leetcode

func CanJump(nums []int) bool {

	var next int = 0
	var can bool = true
	for next < len(nums)-1 {
		max := nums[next]
		if max == 0 {
			can = false
			break
		}

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
	}

	return can
}
