package leetcode

func NextPermutation(nums []int) {

	if len(nums) < 2 {
		return
	}

	var b int = -1
	for r := len(nums) - 1; r > 0; r-- {
		if nums[r-1] < nums[r] {
			b = r - 1
			break
		}
	}

	if b == -1 {
		for j, k := b+1, len(nums)-1; j < k; j, k = j+1, k-1 {
			nums[j], nums[k] = nums[k], nums[j]
		}
		return
	}

	var diff, c int = 1<<31 - 1, -1
	for l := b + 1; l < len(nums); l++ {
		if nums[l] > nums[b] && nums[l]-nums[b] <= diff {
			diff = nums[l] - nums[b]
			c = l
		}
	}

	nums[b], nums[c] = nums[c], nums[b]
	for j, k := b+1, len(nums)-1; j < k; j, k = j+1, k-1 {
		nums[j], nums[k] = nums[k], nums[j]
	}
}
