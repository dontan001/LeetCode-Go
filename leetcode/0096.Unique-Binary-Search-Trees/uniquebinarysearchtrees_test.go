package leetcode

import "testing"

func TestNumTrees(t *testing.T) {

	t.Run("test 1", func(t *testing.T) {
		var i, num int = 1, 0

		for ; i <= 10; i++ {
			num = NumTrees(i)
			t.Logf("num - %v: %v", i, num)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		var i, num int = 1, 0
		for ; i <= 10; i++ {
			num = NumTreesDP(i)
			t.Logf("num dp - %v: %v", i, num)
		}
	})

}
