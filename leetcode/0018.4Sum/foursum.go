package leetcode

import (
	"sort"
)

func FourSum(nums []int, target int) [][]int {

	sort.Ints(nums)

	quadruplets := make([][]int, 0)
	for first := 0; first < len(nums)-3; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		for second := first + 1; second < len(nums)-2; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			var start, end int = second + 1, len(nums) - 1
			for start < end {
				sum := nums[first] + nums[second] + nums[start] + nums[end]
				if sum == target {
					quadruplets = append(quadruplets, []int{nums[first], nums[second], nums[start], nums[end]})
				}

				if sum <= target {
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
	}

	return quadruplets
}
