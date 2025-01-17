package leetcode

import "testing"

func TestSimplifyPath(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		path := "/../"
		res := SimplifyPath(path)
		t.Logf("res: %v", res)
	})

	t.Run("case 2", func(t *testing.T) {
		path := "/../a/.."
		res := SimplifyPath(path)
		t.Logf("res: %v", res)
	})

	t.Run("case 3", func(t *testing.T) {
		path := "//a/../../b/../c//.//"
		res := SimplifyPath(path)
		t.Logf("res: %v", res)
	})
}
