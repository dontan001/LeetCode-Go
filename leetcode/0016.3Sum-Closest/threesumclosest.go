package leetcode

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	var sumClosest int = math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		var begin, end int = i + 1, len(nums) - 1
		for begin < end {
			sum := nums[i] + nums[begin] + nums[end]
			if math.Abs(float64(sum-target)) < math.Abs(float64(sumClosest-target)) {
				sumClosest = sum
			}

			if sum == target {
				return sum
			} else if sum > target {
				end--
			} else {
				begin++
			}
		}
	}

	return sumClosest
}

func ThreeSumClosestBruteForce(nums []int, target int) int {
	var minDiff int = math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				diff := sum - target
				if math.Abs(float64(diff)) < math.Abs(float64(minDiff)) {
					minDiff = diff
				}
			}
		}
	}

	return target + minDiff
}
