package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type Case struct {
		nums   []int
		target int
		expect []int
	}
	cases := []Case{
		{[]int{2, 7, 5, 3}, 10, []int{1, 3}},
		{[]int{2, 9, 8, 5}, 10, []int{0, 2}},
	}

	for _, c := range cases {
		result := TwoSumV1(c.nums, c.target)
		fmt.Printf("V1 Expected %v, Actual %v \n", c.expect, result)
	}

	for _, c := range cases {
		result := TwoSumV2(c.nums, c.target)
		fmt.Printf("V2 Expected %v, Actual %v \n", c.expect, result)
	}
}
