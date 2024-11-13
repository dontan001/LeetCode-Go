package leetcode

import (
	"math"
)

func ThreeSumClosestBruteForce(nums []int, target int) int {
	var minDiff int
	var init bool = true
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				diff := sum - target
				if init {
					minDiff = diff
					init = false
				} else {
					if math.Abs(float64(diff)) < math.Abs(float64(minDiff)) {
						minDiff = diff
					}
				}
			}
		}
	}

	return target + minDiff
}
