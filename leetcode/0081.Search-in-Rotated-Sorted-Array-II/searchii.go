package leetcode

func SearchII(nums []int, target int) bool {

	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return true
		}

		if nums[m] > nums[l] {
			if nums[m] > target && nums[l] <= target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] < nums[r] {
			if nums[m] < target && nums[r] >= target {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if nums[m] == nums[l] {
				l++
			}
			if nums[m] == nums[r] {
				r--
			}
		}
	}

	return false
}
