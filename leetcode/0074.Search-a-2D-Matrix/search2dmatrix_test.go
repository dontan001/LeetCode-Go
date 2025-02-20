package leetcode

import "testing"

func TestSearchMatrix(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {

		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		}
		target := 6

		exist := SearchMatrix(matrix, target)
		t.Logf("found: %t", exist)
	})
}
