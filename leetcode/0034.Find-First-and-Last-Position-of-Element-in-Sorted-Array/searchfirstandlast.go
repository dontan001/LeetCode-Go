package leetcode

func SearchRange(nums []int, target int) []int {

	var search func(int, int) int
	search = func(begin int, end int) int {
		if begin == end {
			if nums[begin] == target {
				return begin
			} else {
				return -1
			}
		}

		if nums[begin] > target || nums[end] < target {
			return -1
		}

		middle := (begin + end) / 2
		if nums[middle] >= target {
			return search(begin, middle)
		} else {
			return search(middle+1, end)
		}
	}

	var first, last int = -1, -1
	if len(nums) == 0 {
		return []int{first, last}
	}

	index := search(0, len(nums)-1)
	if index != -1 {
		first, last = index, index
		for first > 0 {
			if nums[first-1] == nums[index] {
				first--
			} else {
				break
			}
		}

		for last < len(nums)-1 {
			if nums[last+1] == nums[index] {
				last++
			} else {
				break
			}
		}
	}

	return []int{first, last}
}
