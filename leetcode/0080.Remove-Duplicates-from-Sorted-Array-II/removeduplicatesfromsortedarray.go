package leetcode

func RemoveDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	var k, idx, counter int = -1, 1, 1
	for idx < len(nums) {
		if nums[idx] != nums[idx-1] {
			counter = 1
		} else {
			counter++
			if counter > 2 && k == -1 {
				k = idx
			}
		}

		if k < idx && k > -1 {
			if nums[idx] != nums[idx-1] {
				nums[k] = nums[idx]
				k++
			} else {
				if counter <= 2 {
					nums[k] = nums[idx]
					k++
				}
			}
		}

		idx++
	}

	if k == -1 {
		k = len(nums)
	}

	return k
}
