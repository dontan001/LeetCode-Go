package leetcode

func InsertInterval(intervals [][]int, newInterval []int) [][]int {

	res := [][]int{}

	var i int = 0
	for i < len(intervals) {
		if newInterval[0] > intervals[i][1] {
			res = append(res, intervals[i])
			i++
		} else {
			break
		}
	}

	var begin, end int = newInterval[0], newInterval[1]
	for i < len(intervals) {
		if end < intervals[i][0] {
			break
		} else {
			if begin > intervals[i][0] {
				begin = intervals[i][0]
			}
			if end < intervals[i][1] {
				end = intervals[i][1]
			}
			i++
		}
	}
	res = append(res, []int{begin, end})

	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}
