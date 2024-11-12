package leetcode

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)

	triplets := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		var start, end int = i + 1, len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum == 0 {
				triplets = append(triplets, []int{nums[i], nums[start], nums[end]})
			}

			if sum <= 0 {
				start++
				for start < end && nums[start] == nums[start-1] {
					start++
				}
			} else {
				end--
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			}
		}
	}

	return triplets
}
