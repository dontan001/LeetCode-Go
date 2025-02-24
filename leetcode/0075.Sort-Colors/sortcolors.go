package leetcode

func SortColors(nums []int) {

	if len(nums) <= 1 {
		return
	}

	var low, mid, high int = 0, 0, len(nums) - 1
	for mid <= high {
		if nums[mid] == 0 {
			if low != mid {
				nums[low], nums[mid] = nums[mid], nums[low]
			}
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}
