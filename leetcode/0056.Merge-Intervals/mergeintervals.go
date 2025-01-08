package leetcode

import "sort"

func MergeIntervals(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int = [][]int{}
	var begin, end int = -1, -1
	for i := 0; i < len(intervals); i++ {
		if i == 0 {
			begin, end = intervals[i][0], intervals[i][1]
		} else {
			if intervals[i][0] <= end {
				if end < intervals[i][1] {
					end = intervals[i][1]
				}
			} else {
				merged = append(merged, []int{begin, end})
				begin, end = intervals[i][0], intervals[i][1]
			}
		}

		if i == len(intervals)-1 {
			merged = append(merged, []int{begin, end})
		}
	}
	return merged
}
