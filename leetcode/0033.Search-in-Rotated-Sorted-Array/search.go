package leetcode

func Search(nums []int, target int) int {

	var searchPivot func(int, int) int
	searchPivot = func(begin int, end int) int {
		if begin >= end {
			return -1
		}

		middle := (begin + end) / 2
		if nums[middle] > nums[middle+1] {
			return middle
		}

		res := searchPivot(begin, middle)
		if res == -1 {
			res = searchPivot(middle+1, end)
		}
		return res
	}

	var searchBinary func(int, int) int
	searchBinary = func(begin, end int) int {
		if begin == end {
			if nums[begin] == target {
				return begin
			} else {
				return -1
			}
		}

		middle := (begin + end) / 2
		if nums[middle] >= target {
			return searchBinary(begin, middle)
		} else {
			return searchBinary(middle+1, end)
		}
	}

	var res, size int = -1, len(nums)
	pivot := searchPivot(0, size-1)
	if pivot == -1 {
		res = searchBinary(0, size-1)
	} else {
		if nums[0] <= target && target <= nums[pivot] {
			res = searchBinary(0, pivot)
		} else if nums[pivot+1] <= target && target <= nums[size-1] {
			res = searchBinary(pivot+1, size-1)
		}
	}
	return res
}
