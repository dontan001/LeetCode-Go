package leetcode

import "testing"

func TestZigZagConvert(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		s := "abcdef"
		numRows := 2
		t.Logf("the result is %s", ZigZagConvert(s, numRows))
	})

	t.Run("case 2", func(t *testing.T) {
		s := "paypalishiring"
		numRows := 3
		t.Logf("the result is %s", ZigZagConvert(s, numRows))
	})
}
