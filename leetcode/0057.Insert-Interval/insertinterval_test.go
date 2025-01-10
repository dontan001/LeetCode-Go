package leetcode

import "testing"

func TestInsertInterval(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		intervals := [][]int{{1, 5}}
		newInterval := []int{0, 0}

		after := InsertInterval(intervals, newInterval)
		t.Logf("after: %v", after)
	})

	t.Run("case 2", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {6, 9}}
		newInterval := []int{2, 5}

		after := InsertInterval(intervals, newInterval)
		t.Logf("after: %v", after)
	})
}
