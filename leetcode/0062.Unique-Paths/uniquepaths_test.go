package leetcode

import "testing"

func TestUniquePaths(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		var m, n int = 3, 2
		total := UniquePaths(m, n)
		t.Logf("UniquePaths total %d", total)
	})

	t.Run("case 2", func(t *testing.T) {
		var m, n int = 100, 100
		total := UniquePaths(m, n)
		t.Logf("UniquePaths total %d", total)
	})
}
